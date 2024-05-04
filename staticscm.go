package eddnc

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"git.fractalqb.de/fractalqb/daq"
	"git.fractalqb.de/fractalqb/eloc"
)

//go:generate versioner -bno build_no -pkg eddnc ./VERSION ./version.go

type Header struct {
	UploaderID       string
	SoftwareName     string
	SoftwareVersion  string
	GatewayTimestamp time.Time `json:",omitempty"`
	daq.DictAny
}

func (h *Header) Wrap(hdr daq.DictAny) {
	h.UploaderID = daq.Must(hdr.AsString("uploaderID"))
	h.SoftwareName = daq.Must(hdr.AsString("softwareName"))
	h.SoftwareVersion = daq.Must(hdr.AsString("softwareVersion"))
	h.GatewayTimestamp = hdr.TimeOr("gatewayTimestamp", time.Time{})
	h.DictAny = hdr
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

func (m *msg) Wrap(msg daq.DictAny) {
	m.T = daq.Must(msg.AsTime("timestamp"))
}

type atStation struct {
	msg
	StationName string
}

func (m *atStation) Wrap(msg daq.DictAny) {
	m.msg.Wrap(msg)
	m.System = daq.Must(msg.AsString("systemName"))
	m.StationName = daq.Must(msg.AsString("stationName"))
}

type Event struct {
	SchemaRef string
	Header    Header
	Message   Message
}

func (e *Event) Parse(txt []byte) error {
	var gen daq.DictAny
	if err := json.Unmarshal(txt, &gen); err != nil {
		return err
	}
	return e.Wrap(gen)
}

func (e *Event) Wrap(evt daq.DictAny) (err error) {
	defer eloc.RecoverAs(&err)
	e.SchemaRef = daq.Must(evt.AsString("$schemaRef"))
	e.Header.Wrap(daq.Must(evt.AsDictAny("header")))
	switch e.SchemaRef {
	case ScmURLs[Sjournal], ScmURLs[Snavbeaconscan], ScmURLs[Sscanbarycentre]:
		// TODO put non-journal schemas into specific messages
		jm := new(JournalMsg)
		if err = jm.Wrap(daq.Must(evt.AsDictAny("message"))); err != nil {
			return err
		}
		e.Message = jm
	case ScmURLs[Scommodity]:
		cm := new(CommodityMsg)
		if err = cm.Wrap(daq.Must(evt.AsDictAny("message"))); err != nil {
			return err
		}
		e.Message = cm
	case ScmURLs[Sfssdiscoveryscan]:
		fds := new(FSSDiscoScanMsg)
		if err = fds.Wrap(daq.Must(evt.AsDictAny("message"))); err != nil {
			return err
		}
		e.Message = fds
	case ScmURLs[Scodexentry]:
		cdx := new(CodexMsg)
		if err = cdx.Wrap(daq.Must(evt.AsDictAny("message"))); err != nil {
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
	daq.DictAny
}

func (je *JournalMsg) Wrap(msg daq.DictAny) error {
	je.msg.Wrap(msg)
	je.System = daq.Must(msg.AsString("StarSystem"))
	je.SystemAddr = daq.Must(msg.AsInt64("SystemAddress"))
	je.Event = daq.Must(msg.AsString("event"))
	spos := daq.Must(msg.AsSliceAny("StarPos"))
	je.StarPos = [3]float64{
		daq.Must(spos.AsFloat64(0)),
		daq.Must(spos.AsFloat64(1)),
		daq.Must(spos.AsFloat64(2)),
	}
	je.DictAny = msg
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

func (cm *CommodityMsg) Wrap(msg daq.DictAny) error {
	cm.atStation.Wrap(msg)
	cm.MarketID = daq.Must(msg.AsInt64("marketId"))
	cmdts := daq.Must(msg.AsSliceAny("commodities"))
	if l := len(cmdts); cap(cm.Commodities) >= l {
		cm.Commodities = cm.Commodities[:l]
	} else {
		cm.Commodities = make([]*Commodity, l)
	}
	fromIntOrStr := func(obj daq.DictAny, att string) (res int) {
		val := obj[att]
		if res, err := daq.ToInt(val); err == nil {
			return res
		}
		if txt, err := daq.ToString(val); err != nil {
			res, _ = strconv.Atoi(txt)
		}
		return
	}
	for i, e := range cmdts {
		var src daq.DictAny = e.(map[string]any)
		dst := cm.Commodities[i]
		if dst == nil {
			dst = new(Commodity)
			cm.Commodities[i] = dst
		}
		dst.Name = daq.Must(src.AsString("name"))
		dst.MeanPrice = daq.Must(src.AsInt("meanPrice"))
		dst.BuyPrice = daq.Must(src.AsInt("buyPrice"))
		dst.Stock = daq.Must(src.AsInt("stock"))
		dst.StockBracket = fromIntOrStr(src, "stockBracket")
		dst.SellPrice = daq.Must(src.AsInt("sellPrice"))
		dst.Demand = daq.Must(src.AsInt("demand"))
		dst.DemandBracket = fromIntOrStr(src, "demandBracket")
		if arr, _ := src.AsSliceAny("statusFlags"); arr == nil {
			dst.StatusFlags = nil
		} else {
			dst.StatusFlags = make([]string, len(arr))
			for i, v := range arr {
				dst.StatusFlags[i] = daq.Must(daq.ToString(v))
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

func (cdx *CodexMsg) Wrap(msg daq.DictAny) error {
	cdx.msg.Wrap(msg)
	cdx.System = daq.Must(msg.AsString("System"))
	cdx.EntryID = daq.Must(msg.AsInt64("EntryID"))
	cdx.SystemAddress = daq.Must(msg.AsInt64("SystemAddress"))
	coos := daq.Must(msg.AsSliceAny("StarPos"))
	cdx.StarPos[0] = daq.Must(coos.AsFloat32(0))
	cdx.StarPos[1] = daq.Must(coos.AsFloat32(1))
	cdx.StarPos[2] = daq.Must(coos.AsFloat32(2))
	if reg := trim(daq.Must(msg.AsString("Region"))); !strings.HasPrefix(reg, "Codex_RegionName_") {
		return fmt.Errorf("illegal region name format: '%s'", reg)
	} else {
		reg = reg[17:]
		rid, err := strconv.Atoi(reg)
		if err != nil {
			return err
		}
		cdx.Region = int16(rid)
	}
	cdx.Name = trim(daq.Must(msg.AsString("Name")))
	cdx.Category = trim(daq.Must(msg.AsString("Category")))
	cdx.SubCategory = trim(daq.Must(msg.AsString("SubCategory")))
	cdx.BodyID = msg.Int16Or("BodyID", -1)
	if cdx.BodyID >= 0 {
		cdx.BodyName = daq.Must(msg.AsString("BodyName"))
		cdx.Latitude = daq.Must(msg.AsFloat32("Latitude"))
		cdx.Longitude = daq.Must(msg.AsFloat32("Longitude"))
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

func (fds *FSSDiscoScanMsg) Wrap(msg daq.DictAny) error {
	fds.msg.Wrap(msg)
	fds.System = daq.Must(msg.AsString("SystemName"))
	fds.SystemAddress = daq.Must(msg.AsInt64("SystemAddress"))
	coos := daq.Must(msg.AsSliceAny("StarPos"))
	fds.StarPos[0] = daq.Must(coos.AsFloat32(0))
	fds.StarPos[1] = daq.Must(coos.AsFloat32(1))
	fds.StarPos[2] = daq.Must(coos.AsFloat32(2))
	fds.BodyCount = daq.Must(msg.AsInt16("BodyCount"))
	fds.NonBodyCount = daq.Must(msg.AsInt16("NonBodyCount"))
	return nil
}
