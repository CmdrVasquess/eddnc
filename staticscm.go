package eddnc

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
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

func trim(s string) string {
	if s == "" {
		return ""
	}
	if s[0] == '$' {
		s = s[1:]
		if s == "" {
			return ""
		}
	}
	if l := len(s); s[l-1] == ';' {
		s = s[:l-1]
	}
	return s
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
	case ScmURLs[Sjournal], ScmURLs[Snavbeaconscan], ScmURLs[Sscanbarycentre]:
		// TODO put non-journal schemas into specific messages
		jm := new(JournalMsg)
		if err = jm.Wrap(*evt.MObj("message")); err != nil {
			return err
		}
		e.Message = jm
	case ScmURLs[Scommodity]:
		cm := new(CommodityMsg)
		if err = cm.Wrap(*evt.MObj("message")); err != nil {
			return err
		}
		e.Message = cm
	case ScmURLs[Sfssdiscoveryscan]:
		fds := new(FSSDiscoScanMsg)
		if err = fds.Wrap(*evt.MObj("message")); err != nil {
			return err
		}
		e.Message = fds
	case ScmURLs[Scodexentry]:
		cdx := new(CodexMsg)
		if err = cdx.Wrap(*evt.MObj("message")); err != nil {
			return err
		}
		e.Message = cdx
	default:
		if _, ok := ScmMap[e.SchemaRef]; !ok {
			return fmt.Errorf("unknown schema: '%s'", e.SchemaRef)
		}
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

func (je *JournalMsg) Wrap(msg ggja.Obj) error {
	je.msg.Wrap(msg)
	je.System = msg.MStr("StarSystem")
	je.SystemAddr = msg.MInt64("SystemAddress")
	je.Event = msg.MStr("event")
	spos := msg.MArr("StarPos")
	je.StarPos = [3]float64{spos.MF64(0), spos.MF64(1), spos.MF64(2)}
	je.Obj = msg
	return nil
}

type CommodityMsg struct {
	atStation
	MarketID    int64
	Commodities []*Commodity
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

func (cm *CommodityMsg) Wrap(msg ggja.Obj) error {
	cm.atStation.Wrap(msg)
	cm.MarketID = msg.MInt64("marketId")
	cmdts := msg.MArr("commodities")
	if l := len(cmdts.Bare); cap(cm.Commodities) >= l {
		cm.Commodities = cm.Commodities[:l]
	} else {
		cm.Commodities = make([]*Commodity, l)
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
		dst := cm.Commodities[i]
		if dst == nil {
			dst = new(Commodity)
			cm.Commodities[i] = dst
		}
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
	return nil
}

type CodexMsg struct {
	msg
	EntryID       int64
	SystemAddress int64
	StarPos       [3]float32
	Region        int16
	Name          string
	Category      string
	SubCategory   string
	BodyID        int16
	BodyName      string
	Latitude      float32
	Longitude     float32
}

func (cdx *CodexMsg) Wrap(msg ggja.Obj) error {
	cdx.msg.Wrap(msg)
	cdx.System = msg.MStr("System")
	cdx.EntryID = msg.MInt64("EntryID")
	cdx.SystemAddress = msg.MInt64("SystemAddress")
	coos := msg.MArr("StarPos")
	cdx.StarPos[0] = coos.MF32(0)
	cdx.StarPos[1] = coos.MF32(1)
	cdx.StarPos[2] = coos.MF32(2)
	if reg := trim(msg.MStr("Region")); !strings.HasPrefix(reg, "Codex_RegionName_") {
		return fmt.Errorf("illegal region name format: '%s'", reg)
	} else {
		reg = reg[17:]
		rid, err := strconv.Atoi(reg)
		if err != nil {
			return err
		}
		cdx.Region = int16(rid)
	}
	cdx.Name = trim(msg.MStr("Name"))
	cdx.Category = trim(msg.MStr("Category"))
	cdx.SubCategory = trim(msg.MStr("SubCategory"))
	cdx.BodyID = msg.Int16("BodyID", -1)
	if cdx.BodyID >= 0 {
		cdx.BodyName = msg.MStr("BodyName")
		cdx.Latitude = msg.MF32("Latitude")
		cdx.Longitude = msg.MF32("Longitude")
	}
	return nil
}

type FSSDiscoScanMsg struct {
	msg
	SystemAddress int64
	StarPos       [3]float32
	BodyCount     int16
	NonBodyCount  int16
}

func (fds *FSSDiscoScanMsg) Wrap(msg ggja.Obj) error {
	fds.msg.Wrap(msg)
	fds.System = msg.MStr("SystemName")
	fds.SystemAddress = msg.MInt64("SystemAddress")
	coos := msg.MArr("StarPos")
	fds.StarPos[0] = coos.MF32(0)
	fds.StarPos[1] = coos.MF32(1)
	fds.StarPos[2] = coos.MF32(2)
	fds.BodyCount = msg.MInt16("BodyCount")
	fds.NonBodyCount = msg.MInt16("NonBodyCount")
	return nil
}
