// generated with go run gen/main.go
package eddnc

const ScmNo = 18

type ScmID int

func (sid ScmID) Info() ScmInfo { return ScmInfos[sid] }

//go:generate stringer -type ScmID
const (
	Sapproachsettlement  ScmID = 0
	Sblackmarket         ScmID = 1
	Scodexentry          ScmID = 2
	Scommodity           ScmID = 3
	Sdockingdenied       ScmID = 4
	Sdockinggranted      ScmID = 5
	Sfcmaterials_capi    ScmID = 6
	Sfcmaterials_journal ScmID = 7
	Sfssallbodiesfound   ScmID = 8
	Sfssbodysignals      ScmID = 9
	Sfssdiscoveryscan    ScmID = 10
	Sfsssignaldiscovered ScmID = 11
	Sjournal             ScmID = 12
	Snavbeaconscan       ScmID = 13
	Snavroute            ScmID = 14
	Soutfitting          ScmID = 15
	Sscanbarycentre      ScmID = 16
	Sshipyard            ScmID = 17
)

type ScmInfo struct {
	Ref     string
	Topic   string
	Version int
}

var ScmInfos = []ScmInfo{
	{allScms[0:49], allScms[0:48], 1},       // https://eddn.edcd.io/schemas/approachsettlement/1
	{allScms[49:91], allScms[49:90], 1},     // https://eddn.edcd.io/schemas/blackmarket/1
	{allScms[91:132], allScms[91:131], 1},   // https://eddn.edcd.io/schemas/codexentry/1
	{allScms[132:172], allScms[132:171], 3}, // https://eddn.edcd.io/schemas/commodity/3
	{allScms[172:216], allScms[172:215], 1}, // https://eddn.edcd.io/schemas/dockingdenied/1
	{allScms[216:261], allScms[216:260], 1}, // https://eddn.edcd.io/schemas/dockinggranted/1
	{allScms[261:308], allScms[261:307], 1}, // https://eddn.edcd.io/schemas/fcmaterials_capi/1
	{allScms[308:358], allScms[308:357], 1}, // https://eddn.edcd.io/schemas/fcmaterials_journal/1
	{allScms[358:406], allScms[358:405], 1}, // https://eddn.edcd.io/schemas/fssallbodiesfound/1
	{allScms[406:451], allScms[406:450], 1}, // https://eddn.edcd.io/schemas/fssbodysignals/1
	{allScms[451:498], allScms[451:497], 1}, // https://eddn.edcd.io/schemas/fssdiscoveryscan/1
	{allScms[498:548], allScms[498:547], 1}, // https://eddn.edcd.io/schemas/fsssignaldiscovered/1
	{allScms[548:586], allScms[548:585], 1}, // https://eddn.edcd.io/schemas/journal/1
	{allScms[586:630], allScms[586:629], 1}, // https://eddn.edcd.io/schemas/navbeaconscan/1
	{allScms[630:669], allScms[630:668], 1}, // https://eddn.edcd.io/schemas/navroute/1
	{allScms[669:710], allScms[669:709], 2}, // https://eddn.edcd.io/schemas/outfitting/2
	{allScms[710:755], allScms[710:754], 1}, // https://eddn.edcd.io/schemas/scanbarycentre/1
	{allScms[755:794], allScms[755:793], 2}, // https://eddn.edcd.io/schemas/shipyard/2
}

var ScmMap = map[string]ScmID{
	allScms[0:49]:    Sapproachsettlement,
	allScms[49:91]:   Sblackmarket,
	allScms[91:132]:  Scodexentry,
	allScms[132:172]: Scommodity,
	allScms[172:216]: Sdockingdenied,
	allScms[216:261]: Sdockinggranted,
	allScms[261:308]: Sfcmaterials_capi,
	allScms[308:358]: Sfcmaterials_journal,
	allScms[358:406]: Sfssallbodiesfound,
	allScms[406:451]: Sfssbodysignals,
	allScms[451:498]: Sfssdiscoveryscan,
	allScms[498:548]: Sfsssignaldiscovered,
	allScms[548:586]: Sjournal,
	allScms[586:630]: Snavbeaconscan,
	allScms[630:669]: Snavroute,
	allScms[669:710]: Soutfitting,
	allScms[710:755]: Sscanbarycentre,
	allScms[755:794]: Sshipyard,
}

const allScms = "https://eddn.edcd.io/schemas/approachsettlement/1https://eddn.edcd.io/schemas/blackmarket/1https://eddn.edcd.io/schemas/codexentry/1https://eddn.edcd.io/schemas/commodity/3https://eddn.edcd.io/schemas/dockingdenied/1https://eddn.edcd.io/schemas/dockinggranted/1https://eddn.edcd.io/schemas/fcmaterials_capi/1https://eddn.edcd.io/schemas/fcmaterials_journal/1https://eddn.edcd.io/schemas/fssallbodiesfound/1https://eddn.edcd.io/schemas/fssbodysignals/1https://eddn.edcd.io/schemas/fssdiscoveryscan/1https://eddn.edcd.io/schemas/fsssignaldiscovered/1https://eddn.edcd.io/schemas/journal/1https://eddn.edcd.io/schemas/navbeaconscan/1https://eddn.edcd.io/schemas/navroute/1https://eddn.edcd.io/schemas/outfitting/2https://eddn.edcd.io/schemas/scanbarycentre/1https://eddn.edcd.io/schemas/shipyard/2"
