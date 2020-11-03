package eddnc

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"git.fractalqb.de/fractalqb/ggja"
)

//go:generate versioner -bno build_no -pkg eddnc ./VERSION ./version.go

type Header struct {
	UploaderID       string
	SoftwareName     string
	SoftwareVersion  string
	GatewayTimestamp time.Time `json:",omitempty"`
	ggja.Obj
}

func (h *Header) Wrap(hdr ggja.Obj) {
	h.UploaderID = hdr.MStr("uploaderID")
	h.SoftwareName = hdr.MStr("softwareName")
	h.SoftwareVersion = hdr.MStr("softwareVersion")
	h.GatewayTimestamp = hdr.Time("gatewayTimestamp", time.Time{})
	h.Obj = hdr
}

type Message interface {
	Timestamp() time.Time
	SystemName() string
}

type msg struct {
	T      time.Time
	System string
}

func (m *msg) Timestamp() time.Time { return m.T }
func (m *msg) SystemName() string   { return m.System }

func (m *msg) Wrap(msg ggja.Obj) {
	m.T = msg.MTime("timestamp")
}

type atStation struct {
	msg
	StationName string
}

func (m *atStation) Wrap(msg ggja.Obj) {
	m.msg.Wrap(msg)
	m.System = msg.MStr("systemName")
	m.StationName = msg.MStr("stationName")
}

type Event struct {
	SchemaRef string
	Header    Header
	Message   Message
}

func (e *Event) Parse(txt []byte) error {
	gen := make(ggja.BareObj)
	if err := json.Unmarshal(txt, &gen); err != nil {
		return err
	}
	return e.Wrap(ggja.Obj{Bare: gen})
}

func (e *Event) Wrap(evt ggja.Obj) (err error) {
	evt.OnError = func(e error) { err = e }
	e.SchemaRef = evt.MStr("$schemaRef")
	e.Header.Wrap(*evt.MObj("header"))
	switch e.SchemaRef {
	case ScmURLs[Sjournal]:
		jm := new(JournalMsg)
		jm.Wrap(*evt.MObj("message"))
		e.Message = jm
	case ScmURLs[Scommodity]:
		cm := new(CommodityMsg)
		cm.Wrap(*evt.MObj("message"))
		e.Message = cm
	default:
		return fmt.Errorf("unknown schema: '%s'", e.SchemaRef)
	}
	return err
}

// TODO
// blackmarket [ "timestamp", "systemName", "stationName", "name", "sellPrice", "prohibited" ]
// outfit      [ "timestamp", "systemName", "stationName", "marketId", "modules" ]
// shipyard    [ "timestamp", "systemName", "stationName", "marketId", "ships" ]

type JournalMsg struct {
	msg
	SystemAddr int64
	StarPos    [3]float64
	Event      string
	ggja.Obj
}

func (je *JournalMsg) Wrap(msg ggja.Obj) {
	je.msg.Wrap(msg)
	je.System = msg.MStr("StarSystem")
	je.SystemAddr = msg.MInt64("SystemAddress")
	je.Event = msg.MStr("event")
	spos := msg.MArr("StarPos")
	je.StarPos = [3]float64{spos.MF64(0), spos.MF64(1), spos.MF64(2)}
	je.Obj = msg
}

type CommodityMsg struct {
	atStation
	MarketID    int
	Commodities []Commodity
}

type Commodity struct {
	Name          string
	MeanPrice     int
	BuyPrice      int
	Stock         int
	StockBracket  int
	SellPrice     int
	Demand        int
	DemandBracket int
	StatusFlags   []string
}

func (cm *CommodityMsg) Wrap(msg ggja.Obj) {
	cm.atStation.Wrap(msg)
	cm.MarketID = msg.MInt("marketId")
	cmdts := msg.MArr("commodities")
	if l := len(cmdts.Bare); cap(cm.Commodities) >= l {
		cm.Commodities = cm.Commodities[:l]
	} else {
		cm.Commodities = make([]Commodity, l)
	}
	fromIntOrStr := func(obj ggja.Obj, att string) (res int) {
		bak := obj.OnError
		defer func() { obj.OnError = bak }()
		intercepted := false
		obj.OnError = func(err error) {
			var nce ggja.NoConversionError
			if errors.As(err, &nce) {
				if str, ok := nce.Value.(string); ok {
					if str == "" {
						log.Debugf("set empty %s string to 0", att)
						res = 0
						intercepted = true
						return
					}
					if res, err = strconv.Atoi(str); err == nil {
						log.Debugf("%s converted from string '%s'", att, str)
						intercepted = true
						return
					}
				}
			}
			bak(err)
		}
		if i := obj.MInt(att); intercepted {
			return res
		} else {
			return i
		}
	}
	for i, e := range cmdts.Bare {
		src := ggja.Obj{Bare: e.(ggja.BareObj), OnError: msg.OnError}
		dst := &cm.Commodities[i]
		dst.Name = src.MStr("name")
		dst.MeanPrice = src.MInt("meanPrice")
		dst.BuyPrice = src.MInt("buyPrice")
		dst.Stock = src.MInt("stock")
		dst.StockBracket = fromIntOrStr(src, "stockBracket")
		dst.SellPrice = src.MInt("sellPrice")
		dst.Demand = src.MInt("demand")
		dst.DemandBracket = fromIntOrStr(src, "demandBracket")
		if arr := src.Arr("statusFlags"); arr == nil {
			dst.StatusFlags = nil
		} else {
			dst.StatusFlags = make([]string, arr.Len())
			for i, v := range arr.Bare {
				dst.StatusFlags[i] = v.(string)
			}
		}
	}
}
