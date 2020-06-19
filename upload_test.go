package eddn

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

var testSwVersion = fmt.Sprintf("%d.%d.%d", Major, Minor, Patch)

func TestValidate(t *testing.T) {
	u := Upload{Vaildate: true, TestUrl: true, DryRun: !testing.Verbose()}
	u.Http.Timeout = 6 * time.Second
	u.Header.Uploader = "_test_"
	u.Header.SwName = "goEDDNc"
	u.Header.SwVersion = testSwVersion
	msg := make(map[string]interface{})
	err := json.Unmarshal([]byte(`{
    "systemName": "Munfayl",
    "stationName": "Samson",
    "timestamp": "2016-10-01T16:01:18Z",
	"marketId": 4711,
    "ships": [
      "Adder",
      "Asp_Scout",
      "CobraMkIII",
      "Python",
      "SideWinder",
      "Viper"
    ]}`), &msg)
	if err != nil {
		t.Fatal(err)
	}
	err = u.Send(Sshipyard, msg)
	if err != nil {
		t.Error(err)
	}
}

func initUpload(u *Upload, t *testing.T) {
	*u = Upload{Vaildate: true, TestUrl: true, DryRun: !testing.Verbose()}
}

func TestCommodityJ(t *testing.T) {
	var u Upload
	initUpload(&u, t)
	u.Http.Timeout = 6 * time.Second
	u.Header.Uploader = "_test_"
	u.Header.SwName = "goEDDNc"
	u.Header.SwVersion = testSwVersion
	msg := NewMessage(Ts(time.Now()))
	market := make(map[string]interface{})
	marketStr := `{ "timestamp":"2018-07-15T12:28:33Z",
	                "event":"Market", "MarketID":3507400192,
					"StationName":"Maine Observatory", "StarSystem":"Ngandan",
					"Items":[
						{ "id":128049152,
						  "Name":"$platinum_name;",
						  "Name_Localised":"Platinum",
						  "Category":"$MARKET_category_metals;",
						  "Category_Localised":"Metals",
						  "BuyPrice":0, "SellPrice":41794,
						  "MeanPrice":19756,
						  "StockBracket":0, "DemandBracket":3,
						  "Stock":0, "Demand":45,
						  "Consumer":true,
						  "Producer":false,
						  "Rare":false },
						{ "id":128049153,
						  "Name":"$palladium_name;",
						  "Name_Localised":"Palladium",
						  "Category":"$MARKET_category_metals;",
						  "Category_Localised":"Metals",
						  "BuyPrice":0, "SellPrice":13835,
						  "MeanPrice":13244,
						  "StockBracket":0, "DemandBracket":3,
						  "Stock":0, "Demand":62,
						  "Consumer":true,
						  "Producer":false,
						  "Rare":false }
] }`
	err := json.Unmarshal([]byte(marketStr), &market)
	if err != nil {
		t.Fatal(err)
	}
	err = SetCommoditiesJ(msg, market)
	if err != nil {
		t.Fatal(err)
	}
	err = u.Send(Scommodity, msg)
	if err != nil {
		t.Error(err)
	}
}

func TestOutfittingJ(t *testing.T) {
	var u Upload
	initUpload(&u, t)
	u.Http.Timeout = 6 * time.Second
	u.Header.Uploader = "_test_"
	u.Header.SwName = "goEDDNc"
	u.Header.SwVersion = testSwVersion
	msg := NewMessage(Ts(time.Now()))
	market := make(map[string]interface{})
	marketStr := `{ "timestamp":"2018-09-22T11:56:06Z", "event":"Outfitting", "MarketID":3223182848, "StationName":"Jensen Gateway", "StarSystem":"64 Ceti", "Horizons":true, "Items":[
{ "id":128049511, "Name":"hpt_advancedtorppylon_fixed_large", "BuyPrice":134266 },
{ "id":128891602, "Name":"hpt_dumbfiremissilerack_fixed_large", "BuyPrice":868275 },
{ "id":128049509, "Name":"hpt_advancedtorppylon_fixed_small", "BuyPrice":9520 },
{ "id":128666725, "Name":"hpt_dumbfiremissilerack_fixed_medium", "BuyPrice":204340 },
{ "id":128666724, "Name":"hpt_dumbfiremissilerack_fixed_small", "BuyPrice":27349 },
{ "id":128671448, "Name":"hpt_minelauncher_fixed_small_impulse", "BuyPrice":30932 },
{ "id":128049500, "Name":"hpt_minelauncher_fixed_small", "BuyPrice":20621 },
{ "id":128049493, "Name":"hpt_basicmissilerack_fixed_medium", "BuyPrice":435540 },
{ "id":128049492, "Name":"hpt_basicmissilerack_fixed_small", "BuyPrice":61710 },
{ "id":128049489, "Name":"hpt_railgun_fixed_medium", "BuyPrice":350880 },
{ "id":128049488, "Name":"hpt_railgun_fixed_small", "BuyPrice":43860 },
{ "id":128049343, "Name":"python_armour_mirrored", "BuyPrice":103013699 },
{ "id":128049344, "Name":"python_armour_reactive", "BuyPrice":114152932 }
 ] }`
	err := json.Unmarshal([]byte(marketStr), &market)
	if err != nil {
		t.Fatal(err)
	}
	err = SetOutfittingJ(msg, market)
	if err != nil {
		t.Fatal(err)
	}
	err = u.Send(Soutfitting, msg)
	if err != nil {
		t.Error(err)
	}
}

func TestShipyardJ(t *testing.T) {
	var u Upload
	initUpload(&u, t)
	u.Http.Timeout = 6 * time.Second
	u.Header.Uploader = "_test_"
	u.Header.SwName = "goEDDNc"
	u.Header.SwVersion = testSwVersion
	msg := NewMessage(Ts(time.Now()))
	shipyard := make(map[string]interface{})
	shipyardStr := `{ "timestamp":"2018-09-30T18:39:19Z", "event":"Shipyard", "MarketID":3518140160, "StationName":"Koolhaas' Folly", "StarSystem":"HIP 73293", "Horizons":true, "AllowCobraMkIV":false, "PriceList":[
{ "id":128049249, "ShipType":"sidewinder", "ShipPrice":32000 },
{ "id":128049261, "ShipType":"hauler", "ShipPrice":52720 },
{ "id":128672138, "ShipType":"empire_eagle", "ShipType_Localised":"Imperial Eagle", "ShipPrice":110825 },
{ "id":128672255, "ShipType":"viper_mkiv", "ShipType_Localised":"Viper MkIV", "ShipPrice":437931 },
{ "id":128049291, "ShipType":"dolphin", "ShipPrice":1337323 },
{ "id":128049333, "ShipType":"type9", "ShipType_Localised":"Type-9 Heavy", "ShipPrice":76555842 },
{ "id":128049345, "ShipType":"belugaliner", "ShipType_Localised":"Beluga Liner", "ShipPrice":84532764 }
 ] }
`
	err := json.Unmarshal([]byte(shipyardStr), &shipyard)
	if err != nil {
		t.Fatal(err)
	}
	err = SetShipyardJ(msg, shipyard)
	if err != nil {
		t.Fatal(err)
	}
	err = u.Send(Sshipyard, msg)
	if err != nil {
		t.Error(err)
	}
}
