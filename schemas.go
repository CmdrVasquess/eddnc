// generated with genschemas.sh
package eddnc

const ScmNo = 12

//go:generate stringer -type ScmID
const (
	Sapproachsettlement ScmID = iota
	Sblackmarket
	Scodexentry
	Scommodity
	Sfssallbodiesfound
	Sfssdiscoveryscan
	Sjournal
	Snavbeaconscan
	Snavroute
	Soutfitting
	Sscanbarycentre
	Sshipyard
)

var ScmURLs = []string{
	"https://eddn.edcd.io/schemas/approachsettlement/1",
	"https://eddn.edcd.io/schemas/blackmarket/1",
	"https://eddn.edcd.io/schemas/codexentry/1",
	"https://eddn.edcd.io/schemas/commodity/3",
	"https://eddn.edcd.io/schemas/fssallbodiesfound/1",
	"https://eddn.edcd.io/schemas/fssdiscoveryscan/1",
	"https://eddn.edcd.io/schemas/journal/1",
	"https://eddn.edcd.io/schemas/navbeaconscan/1",
	"https://eddn.edcd.io/schemas/navroute/1",
	"https://eddn.edcd.io/schemas/outfitting/2",
	"https://eddn.edcd.io/schemas/scanbarycentre/1",
	"https://eddn.edcd.io/schemas/shipyard/2",
}

var ScmMap = map[string]ScmID{
	"https://eddn.edcd.io/schemas/approachsettlement/1": 0,
	"https://eddn.edcd.io/schemas/blackmarket/1":        1,
	"https://eddn.edcd.io/schemas/codexentry/1":         2,
	"https://eddn.edcd.io/schemas/commodity/3":          3,
	"https://eddn.edcd.io/schemas/fssallbodiesfound/1":  4,
	"https://eddn.edcd.io/schemas/fssdiscoveryscan/1":   5,
	"https://eddn.edcd.io/schemas/journal/1":            6,
	"https://eddn.edcd.io/schemas/navbeaconscan/1":      7,
	"https://eddn.edcd.io/schemas/navroute/1":           8,
	"https://eddn.edcd.io/schemas/outfitting/2":         9,
	"https://eddn.edcd.io/schemas/scanbarycentre/1":     10,
	"https://eddn.edcd.io/schemas/shipyard/2":           11,
}

var ScmDefs = []string{
	`{
    "$schema"               : "http://json-schema.org/draft-04/schema#",
    "id"                    : "https://eddn.edcd.io/schemas/approachsettlement/1#",
    "type"                  : "object",
    "additionalProperties"  : false,
    "required": [ "$schemaRef", "header", "message" ],
    "properties": {
        "$schemaRef": {
            "type"                  : "string"
        },
        "header": {
            "type"                  : "object",
            "additionalProperties"  : true,
            "required"              : [ "uploaderID", "softwareName", "softwareVersion" ],
            "properties"            : {
                "uploaderID": {
                    "type"          : "string"
                },
                "softwareName": {
                    "type"          : "string"
                },
                "softwareVersion": {
                    "type"          : "string"
                },
                "gatewayTimestamp": {
                    "type"          : "string",
                    "format"        : "date-time",
                    "description"   : "Timestamp upon receipt at the gateway. If present, this property will be overwritten by the gateway; submitters are not intended to populate this property."
                }
            }
        },
        "message": {
            "type"                  : "object",
            "description"           : "Contains all properties from the listed events in the client's journal minus the Localised strings and the properties marked below as 'disallowed'",
            "additionalProperties"  : false,
            "required"              : [ "timestamp", "event", "StarSystem", "StarPos", "SystemAddress", "Name", "MarketID", "BodyID", "BodyName", "Latitude", "Longitude" ],
            "properties"            : {
                "timestamp": {
                    "type"          : "string",
                    "format"        : "date-time"
                },
                "event" : {
                    "enum"          : [ "ApproachSettlement" ]
                },
                "horizons": {
                    "type"          : "boolean",
                    "description"   : "Whether the sending Cmdr has a Horizons pass."
                },
                "odyssey": {
                    "type"      : "boolean",
                    "description" : "Whether the sending Cmdr has an Odyssey expansion."
                },
                "StarSystem": {
                    "type"          : "string",
                    "minLength"     : 1,
                    "description"   : "Must be added by the sender"
                },
                "StarPos": {
                    "type"          : "array",
                    "items"         : { "type": "number" },
                    "minItems"      : 3,
                    "maxItems"      : 3,
                    "description"   : "Must be added by the sender"
                },
                "SystemAddress": {
                    "type"          : "integer"
                },
                "Name"             : {
                    "type"          : "string",
                    "description"   : "Name of settlement"
                },
                "MarketID": {
                    "type"          : "integer"
                },		
                "BodyID": {
                    "type"          : "integer"
                },
                "BodyName": {
                    "type"          : "string"
                },	
                "Latitude": {
                    "type"          : "number"
                },
                "Longitude": {
                    "type"          : "number"
                }
            }
        }
    },
    "definitions": {
        "disallowed" : { "not" : { "type": [ "array", "boolean", "integer", "number", "null", "object", "string" ] } }
    }
}
`,
	`{
    "$schema"               : "http://json-schema.org/draft-04/schema#",
    "id"                    : "https://eddn.edcd.io/schemas/blackmarket/1#",
    "type"                  : "object",
    "additionalProperties"  : false,
    "required": [ "$schemaRef", "header", "message" ],
    "properties": {
        "$schemaRef": {
            "type"                  : "string"
        },
        "header": {
            "type"                  : "object",
            "additionalProperties"  : true,
            "required"              : [ "uploaderID", "softwareName", "softwareVersion" ],
            "properties"            : {
                "uploaderID": {
                    "type"          : "string"
                },
                "softwareName": {
                    "type"          : "string"
                },
                "softwareVersion": {
                    "type"          : "string"
                },
                "gatewayTimestamp": {
                    "type"          : "string",
                    "format"        : "date-time",
                    "description"   : "Timestamp upon receipt at the gateway. If present, this property will be overwritten by the gateway; submitters are not intended to populate this property."
                }
            }
        },
        "message": {
            "type"                  : "object",
            "description"           : "Contains all properties from the listed events in the client's journal minus Localised strings and the properties marked below as 'disallowed'",
            "additionalProperties"  : true,
            "required"              : [ "systemName", "stationName", "timestamp", "name", "sellPrice", "prohibited" ],
            "properties"            : {
                "systemName": {
                    "type"          : "string",
                    "minLength"     : 1
                },
                "stationName": {
                    "type"          : "string",
                    "minLength"     : 1
                },                
                "marketId": {
                    "type"          : "integer",
                    "renamed"       : "MarketID"
                },
                "timestamp": {
                    "type"          : "string",
                    "format"        : "date-time"
                },
                "name": {
                    "type"          : "string",
                    "renamed"       : "Type",
                    "minLength"     : 1,
                    "description"   : "Commodity name as returned by the MarketSell entry in the Journal"
                },
                "sellPrice": {
                    "type"          : "integer",
                    "description"   : "Price to sell to the market"
                },
                "prohibited": {
                    "type"          : "boolean",
                    "renamed"       : "IllegalGoods",
                    "description"   : "Whether the commodity is prohibited at this station"
                }
            }
        }
    }
}
`,
	`{
    "$schema"               : "http://json-schema.org/draft-04/schema#",
    "id"                    : "https://eddn.edcd.io/schemas/codexentry/1#",
    "description"           : "EDDN schema for CodexEntry Journal events.  Full documentation at https://github.com/EDCD/EDDN/tree/master/schemas/codexentry-README.md",
    "type"                  : "object",
    "additionalProperties"  : false,
    "required": [ "$schemaRef", "header", "message" ],
    "properties": {
        "$schemaRef": {
            "type"                  : "string"
        },
        "header": {
            "type"                  : "object",
            "additionalProperties"  : true,
            "required"              : [ "uploaderID", "softwareName", "softwareVersion" ],
            "properties"            : {
                "uploaderID": {
                    "type"          : "string"
                },
                "softwareName": {
                    "type"          : "string"
                },
                "softwareVersion": {
                    "type"          : "string"
                },
                "gatewayTimestamp": {
                    "type"          : "string",
                    "format"        : "date-time",
                    "description"   : "Timestamp upon receipt at the gateway. If present, this property will be overwritten by the gateway; submitters are not intended to populate this property."
                }
            }
        },
        "message": {
            "type"                  : "object",
            "description"           : "Contains all properties from the listed events in the client's journal minus Localised strings and the properties marked below as 'disallowed'",
            "additionalProperties"  : false,
            "required"              : [ "timestamp", "event", "System", "StarPos", "SystemAddress", "EntryID" ],
            "properties"            : {
                "timestamp": {
                    "type"          : "string",
                    "format"        : "date-time"
                },
                "event": {
                    "enum"          : [ "CodexEntry" ]
                },
                "horizons": {
                    "type"          : "boolean",
                    "description"   : "Whether the sending Cmdr has a Horizons pass."
                },
                "odyssey": {
                    "type"      : "boolean",
                    "description" : "Whether the sending Cmdr has an Odyssey expansion."
                },
                "System": {
                    "type"          : "string",
                    "minLength"     : 1
                },
                "StarPos": {
                    "type"          : "array",
                    "items"         : { "type": "number" },
                    "minItems"      : 3,
                    "maxItems"      : 3,
                    "description"   : "Must be added by the sender if not present in the journal event"
                },
                "SystemAddress": {
                    "type"          : "integer",
                    "description"   : "Should be added by the sender if not present in the journal event"
                },
                "Name": {
                    "type"          : "string",
                    "minLength"     : 1
                },
                "Region": {
                    "type"          : "string",
                    "minLength"     : 1
                },
                "EntryID": {
                    "type"          : "integer"
                },
                "Category": {
                    "type"          : "string",
                    "minLength"     : 1
                },
                "Latitude": {
                    "type"          : "number"
                },
                "Longitude": {
                    "type"          : "number"
                },
                "SubCategory": {
                    "type"          : "string",
                    "minLength"     : 1
                },
                "NearestDestination": {
                    "type"          : "string"
                },
                "VoucherAmount": {
                    "type"          : "integer"
                },
                "Traits": {
                    "type"          : "array",
                    "items"         : {
                        "type"          : "string",
                        "minLength"     : 1
                    }
                },
                "BodyID": {
                    "type"          : "integer"
                },
                "BodyName": {
                    "type"          : "string"
                },
                "IsNewEntry": {
                    "$ref"          : "#/definitions/disallowed",
                    "description"   : "Contains personal data"
                },
                "NewTraitsDiscovered": {
                    "$ref"          : "#/definitions/disallowed",
                    "description"   : "Contains personal data"
                }
            },
            "patternProperties": {
                "_Localised$"       : { "$ref" : "#/definitions/disallowed" }
            }
        }
    },
    "definitions": {
        "disallowed" : { "not" : { "type": [ "array", "boolean", "integer", "number", "null", "object", "string" ] } }
    }
}
`,
	`{
    "$schema"               : "http://json-schema.org/draft-04/schema#",
    "id"                    : "https://eddn.edcd.io/schemas/commodity/3#",
    "type"                  : "object",
    "additionalProperties"  : false,
    "required": [ "$schemaRef", "header", "message" ],
    "properties": {
        "$schemaRef": {
            "type"                  : "string"
        },
        "header": {
            "type"                  : "object",
            "additionalProperties"  : true,
            "required"              : [ "uploaderID", "softwareName", "softwareVersion" ],
            "properties"            : {
                "uploaderID": {
                    "type"          : "string"
                },
                "softwareName": {
                    "type"          : "string"
                },
                "softwareVersion": {
                    "type"          : "string"
                },
                "gatewayTimestamp": {
                    "type"          : "string",
                    "format"        : "date-time",
                    "description"   : "Timestamp upon receipt at the gateway. If present, this property will be overwritten by the gateway; submitters are not intended to populate this property."
                }
            }
        },
        "message": {
            "type"                  : "object",
            "additionalProperties"  : false,
            "required"              : [ "systemName", "stationName", "marketId", "timestamp", "commodities" ],
            "properties"            : {
                "systemName": {
                    "type"      : "string",
                    "minLength" : 1
                },
                "stationName": {
                    "type"      : "string",
                    "renamed"   : "StarSystem",
                    "minLength" : 1
                },                
                "marketId": {
                    "type"      : "integer",
                    "renamed"   : "MarketID"
                },
                "horizons": {
                    "type"      : "boolean",
                    "description" : "Whether the sending Cmdr has a Horizons pass."
                },
                "odyssey": {
                    "type"      : "boolean",
                    "description" : "Whether the sending Cmdr has an Odyssey expansion."
                },                
                "timestamp": {
                    "type"      : "string",
                    "format"    : "date-time"
                },
                "commodities": {
                    "type"      : "array",
                    "description" : "Commodities returned by the Companion API, with illegal commodities omitted",
                    "items"     : {
                        "type"                  : "object",
                        "additionalProperties"  : false,
                        "required"              : [ "name", "meanPrice", "buyPrice", "stock", "stockBracket", "sellPrice", "demand", "demandBracket" ],
                        "properties"            : {
                            "name": {
                                "type"          : "string",
                                "renamed"       : "Name",
                                "minLength"     : 1,
                                "description"   : "Symbolic name as returned by the Companion API"
                            },
                            "meanPrice": {
                                "type"          : "integer",
                                "renamed"       : "MeanPrice"
                            },
                            "buyPrice": {
                                "type"          : "integer",
                                "renaamed"      : "BuyPrice",
                                "description"   : "Price to buy from the market"
                            },
                            "stock": {
                                "type"          : "integer",
                                "renamed"       : "Stock"
                            },
                            "stockBracket": {
                                "$ref"          : "#/definitions/levelType",
                                "renamed"       : "StockBracket"
                            },
                            "sellPrice": {
                                "type"          : "integer",
                                "renamed"       : "SellPrice",
                                "description"   : "Price to sell to the market"
                            },
                            "demand": {
                                "type"          : "integer",
                                "renamed"       : "Demand"
                            },
                            "demandBracket": {
                                "$ref"          : "#/definitions/levelType",
                                "renamed"       : "DemandBracket"
                            },
                            "statusFlags": {
                                "type"          : "array",
                                "minItems"      : 1,
                                "uniqueItems"   : true,
                                "items"         : {
                                    "type"          : "string",
                                    "minLength"     : 1
                                }
                            },
                            "Producer": {
                                "$ref"          : "#/definitions/disallowed",
                                "description"   : "Not present in CAPI data, so removed from Journal-sourced data"
                            },
                            "Rare" : {
                                "$ref"          : "#/definitions/disallowed",
                                "description"   : "Not present in CAPI data, so removed from Journal-sourced data"
                            },
                            "id": {
                                "$ref"          : "#/definitions/disallowed",
                                "description"   : "Not wanted for historical reasons?"
                            }
                        }
                    }
                },
                "economies": {
                    "type"      : "array",
                    "items"     : {
                        "type"                  : "object",
                        "additionalProperties"  : false,
                        "required"              : [ "name", "proportion" ],
                        "properties"            : {
                            "name": {
                                "type"          : "string",
                                "minLength"     : 1,
                                "description"   : "Economy type as returned by the Companion API"
                            },
                            "proportion": {
                                "type"          : "number"
                            }
                        }
                    }
                },
                "prohibited": {
                    "type"          : "array",
                    "uniqueItems"   : true,
                    "items"         : {
                        "type"          : "string",
                        "minLength"     : 1
                    }
                },
                "StationType": {
                    "$ref"          : "#/definitions/disallowed",
                    "description"   : "Not present in CAPI data, so removed from Journal-sourced data"
                }
            }
        }
    },
    "definitions": {
        "disallowed" : { "not" : { "type": [ "array", "boolean", "integer", "number", "null", "object", "string" ] } },
        "levelType": {
            "enum"          : [0, 1, 2, 3, ""],
            "description"   : "Note: A value of \"\" indicates that the commodity is not normally sold/purchased at this station, but is currently temporarily for sale/purchase"
        }
    }
}
`,
	`{
    "$schema"               : "http://json-schema.org/draft-04/schema#",
    "id"                    : "https://eddn.edcd.io/schemas/fssallbodiesfound/1#",
    "type"                  : "object",
    "additionalProperties"  : false,
    "required": [ "$schemaRef", "header", "message" ],
    "properties": {
        "$schemaRef": {
            "type"                  : "string"
        },
        "header": {
            "type"                  : "object",
            "additionalProperties"  : true,
            "required"              : [ "uploaderID", "softwareName", "softwareVersion" ],
            "properties"            : {
                "uploaderID": {
                    "type"          : "string"
                },
                "softwareName": {
                    "type"          : "string"
                },
                "softwareVersion": {
                    "type"          : "string"
                },
                "gatewayTimestamp": {
                    "type"          : "string",
                    "format"        : "date-time",
                    "description"   : "Timestamp upon receipt at the gateway. If present, this property will be overwritten by the gateway; submitters are not intended to populate this property."
                }
            }
        },
        "message": {
            "type"                  : "object",
            "description"           : "Contains all properties from the listed events in the client's journal, minus the Localised strings and the properties marked below as 'disallowed'",
            "additionalProperties"  : false,
            "required"              : [ "timestamp", "event", "SystemName", "StarPos", "SystemAddress", "Count" ],
            "properties"            : {
                "timestamp": {
                    "type"          : "string",
                    "format"        : "date-time"
                },
                "event" : {
                    "enum"          : [ "FSSAllBodiesFound" ]
                },
                "horizons": {
                    "type"          : "boolean",
                    "description"   : "Whether the sending Cmdr has a Horizons pass."
                },
                "odyssey": {
                    "type"      : "boolean",
                    "description" : "Whether the sending Cmdr has an Odyssey expansion."
                },
                "SystemName": {
                    "type"          : "string",
                    "minLength"     : 1
                },
                "StarPos": {
                    "type"          : "array",
                    "items"         : { "type": "number" },
                    "minItems"      : 3,
                    "maxItems"      : 3,
                    "description"   : "Must be added by the sender if not present in the journal event"
                },
                "SystemAddress": {
                    "type"          : "integer"
                },
                "Count"             : {
                    "type"          : "integer",
                    "description"   : "Number of bodies in this system"
                }
            }
        }
    },
    "definitions": {
        "disallowed" : { "not" : { "type": [ "array", "boolean", "integer", "number", "null", "object", "string" ] } }
    }
}
`,
	`{
    "$schema"               : "http://json-schema.org/draft-04/schema#",
    "id"                    : "https://eddn.edcd.io/schemas/fssdiscoveryscan/1#",
    "type"                  : "object",
    "additionalProperties"  : false,
    "required": [ "$schemaRef", "header", "message" ],
    "properties": {
        "$schemaRef": {
            "type"                  : "string"
        },
        "header": {
            "type"                  : "object",
            "additionalProperties"  : true,
            "required"              : [ "uploaderID", "softwareName", "softwareVersion" ],
            "properties"            : {
                "uploaderID": {
                    "type"          : "string"
                },
                "softwareName": {
                    "type"          : "string"
                },
                "softwareVersion": {
                    "type"          : "string"
                },
                "gatewayTimestamp": {
                    "type"          : "string",
                    "format"        : "date-time",
                    "description"   : "Timestamp upon receipt at the gateway. If present, this property will be overwritten by the gateway; submitters are not intended to populate this property."
                }
            }
        },
        "message": {
            "type"                  : "object",
            "description"           : "Contains all properties from the listed events in the client's journal minus Localised strings and the properties marked below as 'disallowed'",
            "additionalProperties"  : false,
            "required"              : [ "timestamp", "event", "SystemName", "StarPos", "SystemAddress", "BodyCount", "NonBodyCount" ],
            "properties"            : {
                "timestamp": {
                    "type"          : "string",
                    "format"        : "date-time"
                },
                "event" : {
                    "enum"          : [ "FSSDiscoveryScan" ]
                },
                "horizons": {
                    "type"          : "boolean",
                    "description"   : "Whether the sending Cmdr has a Horizons pass."
                },
                "odyssey": {
                    "type"      : "boolean",
                    "description" : "Whether the sending Cmdr has an Odyssey expansion."
                },
                "SystemName": {
                    "type"          : "string",
                    "minLength"     : 1
                },
                "StarPos": {
                    "type"          : "array",
                    "items"         : { "type": "number" },
                    "minItems"      : 3,
                    "maxItems"      : 3,
                    "description"   : "Must be added by the sender if not present in the journal event"
                },
                "SystemAddress": {
                    "type"          : "integer"
                },
                "Progress"            : {
                    "$ref" : "#/definitions/disallowed",
                    "description": "Contains personal data"
                },
                "BodyCount"           : {
                    "type"          : "integer"
                },
                "NonBodyCount"           : {
                    "type"          : "integer"
                }
            }
        }
    },
    "definitions": {
        "disallowed" : { "not" : { "type": [ "array", "boolean", "integer", "number", "null", "object", "string" ] } }
    }
}
`,
	`{
    "$schema"               : "http://json-schema.org/draft-04/schema#",
    "id"                    : "https://eddn.edcd.io/schemas/journal/1#",
    "type"                  : "object",
    "additionalProperties"  : false,
    "required": [ "$schemaRef", "header", "message" ],
    "properties": {
        "$schemaRef": {
            "type"                  : "string"
        },
        "header": {
            "type"                  : "object",
            "additionalProperties"  : true,
            "required"              : [ "uploaderID", "softwareName", "softwareVersion" ],
            "properties"            : {
                "uploaderID": {
                    "type"          : "string"
                },
                "softwareName": {
                    "type"          : "string"
                },
                "softwareVersion": {
                    "type"          : "string"
                },
                "gatewayTimestamp": {
                    "type"          : "string",
                    "format"        : "date-time",
                    "description"   : "Timestamp upon receipt at the gateway. If present, this property will be overwritten by the gateway; submitters are not intended to populate this property."
                }
            }
        },
        "message": {
            "type"                  : "object",
            "description"           : "Contains all properties from the listed events in the client's journal minus Localised strings and the properties marked below as 'disallowed'",
            "additionalProperties"  : true,
            "required"              : [ "timestamp", "event", "StarSystem", "StarPos", "SystemAddress" ],
            "properties"            : {
                "timestamp": {
                    "type"          : "string",
                    "format"        : "date-time"
                },
                "event" : {
                    "enum"          : [ "Docked", "FSDJump", "Scan", "Location", "SAASignalsFound", "CarrierJump", "CodexEntry" ]
                },
                "horizons": {
                    "type"          : "boolean",
                    "description"   : "Whether the sending Cmdr has a Horizons pass."
                },
                "odyssey": {
                    "type"      : "boolean",
                    "description" : "Whether the sending Cmdr has an Odyssey expansion."
                },                 
                "StarSystem": {
                    "type"          : "string",
                    "minLength"     : 1,
                    "description"   : "Must be added by the sender if not present in the journal event"
                },
                "StarPos": {
                    "type"          : "array",
                    "items"         : { "type": "number" },
                    "minItems"      : 3,
                    "maxItems"      : 3,
                    "description"   : "Must be added by the sender if not present in the journal event"
                },
                "SystemAddress": {
                    "type"          : "integer",
                    "description"   : "Should be added by the sender if not present in the journal event"
                },

                "Factions": {
                    "type"          : "array",
                    "description"   : "Present in Location, FSDJump and CarrierJump messages",
                    "items" : {
                        "type"      : "object",
                        "properties": {
                            "HappiestSystem"  : { "$ref" : "#/definitions/disallowed" },
                            "HomeSystem"      : { "$ref" : "#/definitions/disallowed" },
                            "MyReputation"    : { "$ref" : "#/definitions/disallowed" },
                            "SquadronFaction" : { "$ref" : "#/definitions/disallowed" }
                        },
                        "patternProperties"   : {
                            "_Localised$"     : { "$ref" : "#/definitions/disallowed" }
                        }
                    }
                },
                
                "ActiveFine"          : { "$ref" : "#/definitions/disallowed" },
                "CockpitBreach"       : { "$ref" : "#/definitions/disallowed" },
                "BoostUsed"           : { "$ref" : "#/definitions/disallowed" },
                "FuelLevel"           : { "$ref" : "#/definitions/disallowed" },
                "FuelUsed"            : { "$ref" : "#/definitions/disallowed" },
                "JumpDist"            : { "$ref" : "#/definitions/disallowed" },
                "Latitude"            : { "$ref" : "#/definitions/disallowed" },
                "Longitude"           : { "$ref" : "#/definitions/disallowed" },
                "Wanted"              : { "$ref" : "#/definitions/disallowed" },
                "IsNewEntry"          : { "$ref" : "#/definitions/disallowed" },
                "NewTraitsDiscovered" : { "$ref" : "#/definitions/disallowed" },
                "Traits"              : { "$ref" : "#/definitions/disallowed" },
                "VoucherAmount"       : { "$ref" : "#/definitions/disallowed" }
            },
            "patternProperties"     : {
                "_Localised$"       : { "$ref" : "#/definitions/disallowed" },
                "^(Materials|StationEconomies|Signals)$" : {
                    "type"          : "array",
                    "description"   : "Present in Scan, Docked and SAASignalsFound messages",
                    "items" : {
                        "type"      : "object",
                        "patternProperties"     : {
                            "_Localised$"       : { "$ref" : "#/definitions/disallowed" }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "disallowed" : { "not" : { "type": [ "array", "boolean", "integer", "number", "null", "object", "string" ] } }
    }
}
`,
	`{
    "$schema"               : "http://json-schema.org/draft-04/schema#",
    "id"                    : "https://eddn.edcd.io/schemas/navbeaconscan/1#",
    "type"                  : "object",
    "additionalProperties"  : false,
    "required": [ "$schemaRef", "header", "message" ],
    "properties": {
        "$schemaRef": {
            "type"                  : "string"
        },
        "header": {
            "type"                  : "object",
            "additionalProperties"  : true,
            "required"              : [ "uploaderID", "softwareName", "softwareVersion" ],
            "properties"            : {
                "uploaderID": {
                    "type"          : "string"
                },
                "softwareName": {
                    "type"          : "string"
                },
                "softwareVersion": {
                    "type"          : "string"
                },
                "gatewayTimestamp": {
                    "type"          : "string",
                    "format"        : "date-time",
                    "description"   : "Timestamp upon receipt at the gateway. If present, this property will be overwritten by the gateway; submitters are not intended to populate this property."
                }
            }
        },
        "message": {
            "type"                  : "object",
            "description"           : "Contains all properties from the listed events in the client's journal minus Localised strings and the properties marked below as 'disallowed'",
            "additionalProperties"  : false,
            "required"              : [ "timestamp", "event", "StarSystem", "StarPos", "SystemAddress", "NumBodies" ],
            "properties"            : {
                "timestamp": {
                    "type"          : "string",
                    "format"        : "date-time"
                },
                "event" : {
                    "enum"          : [ "NavBeaconScan" ]
                },
                "horizons": {
                    "type"          : "boolean",
                    "description"   : "Whether the sending Cmdr has a Horizons pass."
                },
                "odyssey": {
                    "type"      : "boolean",
                    "description" : "Whether the sending Cmdr has an Odyssey expansion."
                },
                "StarSystem": {
                    "type"          : "string",
                    "minLength"     : 1,
                    "description"   : "Should be added by the sender if not present in the journal event"
                },
                "StarPos": {
                    "type"          : "array",
                    "items"         : { "type": "number" },
                    "minItems"      : 3,
                    "maxItems"      : 3,
                    "description"   : "Must be added by the sender if not present in the journal event"
                },
                "SystemAddress": {
                    "type"          : "integer"
                },
                "NumBodies"         : {
                    "type"          : "integer"
                }
            }
        }
    },
    "definitions": {
        "disallowed" : { "not" : { "type": [ "array", "boolean", "integer", "number", "null", "object", "string" ] } }
    }
}
`,
	`{
    "$schema"               : "http://json-schema.org/draft-04/schema#",
    "id"                    : "https://eddn.edcd.io/schemas/navroute/1#",
    "type"                  : "object",
    "additionalProperties"  : false,
    "required": [ "$schemaRef", "header", "message" ],
    "properties": {
        "$schemaRef": {
            "type"                  : "string"
        },
        "header": {
            "type"                  : "object",
            "additionalProperties"  : true,
            "required"              : [ "uploaderID", "softwareName", "softwareVersion" ],
            "properties"            : {
                "uploaderID": {
                    "type"          : "string"
                },
                "softwareName": {
                    "type"          : "string"
                },
                "softwareVersion": {
                    "type"          : "string"
                },
                "gatewayTimestamp": {
                    "type"          : "string",
                    "format"        : "date-time",
                    "description"   : "Timestamp upon receipt at the gateway. If present, this property will be overwritten by the gateway; submitters are not intended to populate this property."
                }
            }
        },
        "message": {
            "type"                  : "object",
            "additionalProperties"  : false,
            "required"              : [ "timestamp", "event", "Route" ],
            "properties"            : {
                "timestamp": {
                    "type"          : "string",
                    "format"        : "date-time"
                },
                "event": {
                    "enum"      : [ "NavRoute" ]
                },
                "horizons": {
                    "type"      : "boolean",
                    "description" : "Whether the sending Cmdr has a Horizons pass."
                },
                "odyssey": {
                    "type"      : "boolean",
                    "description" : "Whether the sending Cmdr has an Odyssey expansion."
                },
                "Route": {
                    "type"      : "array",
                    "description" : "Route generated by in game plotter",
                    "items"     : {
                        "type"                  : "object",
                        "additionalProperties"  : false,
                        "required"              : [ "StarSystem", "SystemAddress", "StarPos", "StarClass" ],
                        "properties"            : {
                            "StarSystem": {
                                "type"          : "string",
                                "minLength"     : 1
                            },
                            "StarPos": {
                                "type"          : "array",
                                "items"         : { "type": "number" },
                                "minItems"      : 3,
                                "maxItems"      : 3
                            },
                            "SystemAddress": {
                                "type"          : "integer"
                            },
                            "StarClass": {
                                "type"          : "string",
                                "minLength"     : 1
                            }
                        }
                    }
                }
            }
        }
    }
}
`,
	`{
    "$schema"               : "http://json-schema.org/draft-04/schema#",
    "id"                    : "https://eddn.edcd.io/schemas/outfitting/2#",
    "type"                  : "object",
    "additionalProperties"  : false,
    "required"              : [ "$schemaRef", "header", "message" ],
    "properties"            : {
        "$schemaRef": {
            "type"                  : "string"
        },
        "header": {
            "type"                  : "object",
            "additionalProperties"  : true,
            "required"              : [ "uploaderID", "softwareName", "softwareVersion" ],
            "properties"            : {
                "uploaderID": {
                    "type"          : "string"
                },
                "softwareName": {
                    "type"          : "string"
                },
                "softwareVersion": {
                    "type"          : "string"
                },
                "gatewayTimestamp": {
                    "type"          : "string",
                    "format"        : "date-time",
                    "description"   : "Timestamp upon receipt at the gateway. If present, this property will be overwritten by the gateway; submitters are not intended to populate this property."
                }
            }
        },
        "message": {
            "type"                  : "object",
            "additionalProperties"  : false,
            "required"              : [ "systemName", "stationName", "marketId", "timestamp", "modules" ],
            "properties"            : {
                "systemName": {
                    "type"      : "string",
                    "renamed"   : "StarSystem",
                    "minLength" : 1
                },
                "stationName": {
                    "type"      : "string",
                    "renamed"   : "StationName",
                    "minLength" : 1
                },                
                "marketId": {
                    "type"      : "integer",
                    "renamed"   : "MarketID"
                },
                "horizons": {
                    "type"      : "boolean",
                    "description" : "Whether the sending Cmdr has a Horizons pass."
                },
                "odyssey": {
                    "type"      : "boolean",
                    "description" : "Whether the sending Cmdr has an Odyssey expansion."
                },                 
                "timestamp": {
                    "type"      : "string",
                    "format"    : "date-time"
                },
                "modules": {
                    "type"          : "array",
                    "renamed"       : "Items",
                    "minItems"      : 1,
                    "uniqueItems"   : true,
                    "items"         : {
                        "type"          : "string",
                        "minLength"     : 1,
                        "pattern"       : "(^Hpt_|^hpt_|^Int_|^int_|_Armour_|_armour_)",
                        "description"   : "Module symbolic name. e.g. Hpt_ChaffLauncher_Tiny, Int_Engine_Size3_Class5_Fast, Independant_Trader_Armour_Grade1, etc. Modules that depend on the Cmdr's purchases (e.g. bobbleheads, paintjobs) or rank (e.g. decals and PowerPlay faction-specific modules) should be omitted."
                    }
                }
            }
        }
    }
}
`,
	`{
    "$schema"               : "http://json-schema.org/draft-04/schema#",
    "id"                    : "https://eddn.edcd.io/schemas/scanbarycentre/1#",
    "type"                  : "object",
    "additionalProperties"  : false,
    "required": [ "$schemaRef", "header", "message" ],
    "properties": {
        "$schemaRef": {
            "type"                  : "string"
        },
        "header": {
            "type"                  : "object",
            "additionalProperties"  : true,
            "required"              : [ "uploaderID", "softwareName", "softwareVersion" ],
            "properties"            : {
                "uploaderID": {
                    "type"          : "string"
                },
                "softwareName": {
                    "type"          : "string"
                },
                "softwareVersion": {
                    "type"          : "string"
                },
                "gatewayTimestamp": {
                    "type"          : "string",
                    "format"        : "date-time",
                    "description"   : "Timestamp upon receipt at the gateway. If present, this property will be overwritten by the gateway; submitters are not intended to populate this property."
                }
            }
        },
        "message": {
            "type"                  : "object",
            "description"           : "Contains all properties from the listed events in the client's journal minus Localised strings and the properties marked below as 'disallowed'",
            "additionalProperties"  : false,
            "required"              : [ "timestamp", "event", "StarSystem", "StarPos", "SystemAddress", "BodyID" ],
            "properties"            : {
                "timestamp": {
                    "type"          : "string",
                    "format"        : "date-time"
                },
                "event" : {
                    "enum"          : [ "ScanBaryCentre" ]
                },
                "horizons": {
                    "type"          : "boolean",
                    "description"   : "Whether the sending Cmdr has a Horizons pass."
                },
                "odyssey": {
                    "type"      : "boolean",
                    "description" : "Whether the sending Cmdr has an Odyssey expansion."
                },
                "StarSystem": {
                    "type"          : "string",
                    "minLength"     : 1
                },
                "StarPos": {
                    "type"          : "array",
                    "items"         : { "type": "number" },
                    "minItems"      : 3,
                    "maxItems"      : 3,
                    "description"   : "Must be added by the sender if not present in the journal event"
                },
                "SystemAddress": {
                    "type"          : "integer"
                },
                "BodyID": {
                    "type"          : "integer"
                },
                "SemiMajorAxis": {
                    "type"          : "number"
                },
                "Eccentricity": {
                    "type"          : "number"
                },
                "OrbitalInclination": {
                    "type"          : "number"
                },
                "Periapsis": {
                    "type"          : "number"
                },
                "OrbitalPeriod": {
                    "type"          : "number"
                },
                "AscendingNode": {
                    "type"          : "number"
                },
                "MeanAnomaly": {
                    "type"          : "number"
                }
            }
        }
    }
}
`,
	`{
    "$schema"               : "http://json-schema.org/draft-04/schema#",
    "id"                    : "https://eddn.edcd.io/schemas/shipyard/2#",
    "type"                  : "object",
    "additionalProperties"  : false,
    "required"              : [ "$schemaRef", "header", "message" ],
    "properties"            : {
        "$schemaRef": {
            "type"                  : "string"
        },
        "header": {
            "type"                  : "object",
            "additionalProperties"  : true,
            "required"              : [ "uploaderID", "softwareName", "softwareVersion" ],
            "properties"            : {
                "uploaderID": {
                    "type"          : "string"
                },
                "softwareName": {
                    "type"          : "string"
                },
                "softwareVersion": {
                    "type"          : "string"
                },
                "gatewayTimestamp": {
                    "type"          : "string",
                    "format"        : "date-time",
                    "description"   : "Timestamp upon receipt at the gateway. If present, this property will be overwritten by the gateway; submitters are not intended to populate this property."
                }
            }
        },
        "message": {
            "type"                  : "object",
            "additionalProperties"  : false,
            "required"              : [ "systemName", "stationName", "marketId", "timestamp", "ships" ],
            "properties"            : {
                "systemName": {
                    "type"          : "string",
                    "renamed"       : "StarSystem",
                    "minLength"     : 1
                },
                "stationName": {
                    "type"          : "string",
                    "renamed"       : "StationName",
                    "minLength"     : 1
                },                
                "marketId": {
                    "type"          : "integer",
                    "renamed"       : "MarketID"
                },
                "horizons": {
                    "type"          : "boolean",
                    "description"   : "Whether the sending Cmdr has a Horizons pass."
                },
                "odyssey": {
                    "type"      : "boolean",
                    "description" : "Whether the sending Cmdr has an Odyssey expansion."
                },                 
                "allowCobraMkIV": {
                    "type"          : "boolean",
                    "description"   : "Whether the sending Cmdr can purchase the Cobra MkIV or not."
                },
                "timestamp": {
                    "type"          : "string",
                    "format"        : "date-time"
                },
                "ships": {
                    "type"          : "array",
                    "renamed"       : "PriceList",
                    "minItems"      : 1,
                    "uniqueItems"   : true,
                    "items"         : {
                        "type"          : "string",
                        "minLength"     : 1,
                        "description"   : "Ship symbolic name. i.e. one of: SideWinder, Adder, Anaconda, Asp, Asp_Scout CobraMkIII, CobraMkIV, Cutter, DiamondBack, DiamondBackXL, Eagle, Empire_Courier, Empire_Eagle, Empire_Trader, Federation_Corvette, Federation_Dropship, Federation_Dropship_MkII, Federation_Gunship, FerDeLance, Hauler, Independant_Trader, Orca, Python, Type6, Type7, Type9, Viper, Viper_MkIV, Vulture"
                    }
                }
            }
        }
    }
}
`,
}
