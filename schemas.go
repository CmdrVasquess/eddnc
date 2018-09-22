// generated with genschemas.sh
package eddn

const (
	Sblackmarket ScmId = iota
	Scommodity
	Sjournal
	Soutfitting
	Sshipyard
)

var ScmURLs = []string{
	"https://eddn.edcd.io/schemas/blackmarket/1",
	"https://eddn.edcd.io/schemas/commodity/3",
	"https://eddn.edcd.io/schemas/journal/1",
	"https://eddn.edcd.io/schemas/outfitting/2",
	"https://eddn.edcd.io/schemas/shipyard/2",
}
var ScmDefs = []string{
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
                    "type"          : "number"
                },
                "timestamp": {
                    "type"          : "string",
                    "format"        : "date-time"
                },
                "name": {
                    "type"          : "string",
                    "minLength"     : 1,
                    "description"   : "Commodity name as returned by the MarketSell entry in the Journal"
                },
                "sellPrice": {
                    "type"          : "integer",
                    "description"   : "Price to sell to the market"
                },
                "prohibited": {
                    "type"          : "boolean",
                    "description"   : "Whether the commodity is prohibited at this station"
                }
            }
        }
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
            "required"              : [ "systemName", "stationName", "timestamp", "commodities" ],
            "properties"            : {
                "systemName": {
                    "type"      : "string",
                    "minLength" : 1
                },
                "stationName": {
                    "type"      : "string",
                    "minLength" : 1
                },                
                "marketId": {
                    "type"          : "number"
                },
                "timestamp": {
                    "type"      : "string",
                    "format"    : "date-time"
                },
                "commodities": {
                    "type"      : "array",
                    "minItems"  : 1,
                    "description" : "Commodities returned by the Companion API, with illegal commodities omitted",
                    "items"     : {
                        "type"                  : "object",
                        "additionalProperties"  : false,
                        "required"              : [ "name", "meanPrice", "buyPrice", "stock", "stockBracket", "sellPrice", "demand", "demandBracket" ],
                        "properties"            : {
                            "name": {
                                "type"          : "string",
                                "minLength"     : 1,
                                "description"   : "Symbolic name as returned by the Companion API"
                            },
                            "meanPrice": {
                                "type"          : "integer"
                            },
                            "buyPrice": {
                                "type"          : "integer",
                                "description"   : "Price to buy from the market"
                            },
                            "stock": {
                                "type"          : "integer"
                            },
                            "stockBracket": {
                                "$ref"          : "#/definitions/levelType"
                            },
                            "sellPrice": {
                                "type"          : "integer",
                                "description"   : "Price to sell to the market"
                            },
                            "demand": {
                                "type"          : "integer"
                            },
                            "demandBracket": {
                                "$ref"          : "#/definitions/levelType"
                            },
                            "statusFlags": {
                                "type"          : "array",
                                "minItems"      : 1,
                                "uniqueItems"   : true,
                                "items"         : {
                                    "type"          : "string",
                                    "minLength"     : 1
                                }
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
                }
            }
        }
    },
    "definitions": {
        "levelType": {
            "enum"          : [0, 1, 2, 3, ""],
            "description"   : "Note: A value of \"\" indicates that the commodity is not normally sold/purchased at this station, but is currently temporarily for sale/purchase"
        }
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
            "required"              : [ "timestamp", "event", "StarSystem", "StarPos" ],
            "properties"            : {
                "timestamp": {
                    "type"          : "string",
                    "format"        : "date-time"
                },
                "event" : {
                    "enum"          : [ "Docked", "FSDJump", "Scan", "Location" ]
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
                    "type"          : "number",
                    "description"   : "Should be added by the sender if not present in the journal event"
                },

                "CockpitBreach"     : { "$ref" : "#/definitions/disallowed" },
                "BoostUsed"         : { "$ref" : "#/definitions/disallowed" },
                "FuelLevel"         : { "$ref" : "#/definitions/disallowed" },
                "FuelUsed"          : { "$ref" : "#/definitions/disallowed" },
                "JumpDist"          : { "$ref" : "#/definitions/disallowed" },
                "Latitude"          : { "$ref" : "#/definitions/disallowed" },
                "Longitude"         : { "$ref" : "#/definitions/disallowed" }
            },
            "patternProperties"     : {
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
            "required"              : [ "systemName", "stationName", "timestamp", "modules" ],
            "properties"            : {
                "systemName": {
                    "type"      : "string",
                    "minLength" : 1
                },
                "stationName": {
                    "type"      : "string",
                    "minLength" : 1
                },                
                "marketId": {
                    "type"          : "number"
                },
                "timestamp": {
                    "type"      : "string",
                    "format"    : "date-time"
                },
                "modules": {
                    "type"          : "array",
                    "minItems"      : 1,
                    "uniqueItems"   : true,
                    "items"         : {
                        "type"          : "string",
                        "minLength"     : 1,
                        "pattern"       : "(^Hpt_|^Int_|_Armour_)",
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
            "required"              : [ "systemName", "stationName", "timestamp", "ships" ],
            "properties"            : {
                "systemName": {
                    "type"      : "string",
                    "minLength" : 1
                },
                "stationName": {
                    "type"      : "string",
                    "minLength" : 1
                },                
                "marketId": {
                    "type"          : "number"
                },
                "timestamp": {
                    "type"      : "string",
                    "format"    : "date-time"
                },
                "ships": {
                    "type"          : "array",
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
