// generated with go run gen/main.go
package eddnc

import _ "embed"

const ScmNo = 14

//go:generate stringer -type ScmID
const (
	Sapproachsettlement  ScmID = 0
	Sblackmarket         ScmID = 1
	Scodexentry          ScmID = 2
	Scommodity           ScmID = 3
	Sfssallbodiesfound   ScmID = 4
	Sfssbodysignals      ScmID = 5
	Sfssdiscoveryscan    ScmID = 6
	Sfsssignaldiscovered ScmID = 7
	Sjournal             ScmID = 8
	Snavbeaconscan       ScmID = 9
	Snavroute            ScmID = 10
	Soutfitting          ScmID = 11
	Sscanbarycentre      ScmID = 12
	Sshipyard            ScmID = 13
)

var ScmURLs = []string{
	"https://eddn.edcd.io/schemas/approachsettlement/1",
	"https://eddn.edcd.io/schemas/blackmarket/1",
	"https://eddn.edcd.io/schemas/codexentry/1",
	"https://eddn.edcd.io/schemas/commodity/3",
	"https://eddn.edcd.io/schemas/fssallbodiesfound/1",
	"https://eddn.edcd.io/schemas/fssbodysignals/1",
	"https://eddn.edcd.io/schemas/fssdiscoveryscan/1",
	"https://eddn.edcd.io/schemas/fsssignaldiscovered/1",
	"https://eddn.edcd.io/schemas/journal/1",
	"https://eddn.edcd.io/schemas/navbeaconscan/1",
	"https://eddn.edcd.io/schemas/navroute/1",
	"https://eddn.edcd.io/schemas/outfitting/2",
	"https://eddn.edcd.io/schemas/scanbarycentre/1",
	"https://eddn.edcd.io/schemas/shipyard/2",
}

var ScmMap = map[string]ScmID{
	"https://eddn.edcd.io/schemas/approachsettlement/1":  Sapproachsettlement,
	"https://eddn.edcd.io/schemas/blackmarket/1":         Sblackmarket,
	"https://eddn.edcd.io/schemas/codexentry/1":          Scodexentry,
	"https://eddn.edcd.io/schemas/commodity/3":           Scommodity,
	"https://eddn.edcd.io/schemas/fssallbodiesfound/1":   Sfssallbodiesfound,
	"https://eddn.edcd.io/schemas/fssbodysignals/1":      Sfssbodysignals,
	"https://eddn.edcd.io/schemas/fssdiscoveryscan/1":    Sfssdiscoveryscan,
	"https://eddn.edcd.io/schemas/fsssignaldiscovered/1": Sfsssignaldiscovered,
	"https://eddn.edcd.io/schemas/journal/1":             Sjournal,
	"https://eddn.edcd.io/schemas/navbeaconscan/1":       Snavbeaconscan,
	"https://eddn.edcd.io/schemas/navroute/1":            Snavroute,
	"https://eddn.edcd.io/schemas/outfitting/2":          Soutfitting,
	"https://eddn.edcd.io/schemas/scanbarycentre/1":      Sscanbarycentre,
	"https://eddn.edcd.io/schemas/shipyard/2":            Sshipyard,
}

var ScmDefs = []string{
	approachsettlementSchema,
	blackmarketSchema,
	codexentrySchema,
	commoditySchema,
	fssallbodiesfoundSchema,
	fssbodysignalsSchema,
	fssdiscoveryscanSchema,
	fsssignaldiscoveredSchema,
	journalSchema,
	navbeaconscanSchema,
	navrouteSchema,
	outfittingSchema,
	scanbarycentreSchema,
	shipyardSchema,
}

var (
	//go:embed schemas/approachsettlement-v1.0.json
	approachsettlementSchema string

	//go:embed schemas/blackmarket-v1.0.json
	blackmarketSchema string

	//go:embed schemas/codexentry-v1.0.json
	codexentrySchema string

	//go:embed schemas/commodity-v3.0.json
	commoditySchema string

	//go:embed schemas/fssallbodiesfound-v1.0.json
	fssallbodiesfoundSchema string

	//go:embed schemas/fssbodysignals-v1.0.json
	fssbodysignalsSchema string

	//go:embed schemas/fssdiscoveryscan-v1.0.json
	fssdiscoveryscanSchema string

	//go:embed schemas/fsssignaldiscovered-v1.0.json
	fsssignaldiscoveredSchema string

	//go:embed schemas/journal-v1.0.json
	journalSchema string

	//go:embed schemas/navbeaconscan-v1.0.json
	navbeaconscanSchema string

	//go:embed schemas/navroute-v1.0.json
	navrouteSchema string

	//go:embed schemas/outfitting-v2.0.json
	outfittingSchema string

	//go:embed schemas/scanbarycentre-v1.0.json
	scanbarycentreSchema string

	//go:embed schemas/shipyard-v2.0.json
	shipyardSchema string
)
