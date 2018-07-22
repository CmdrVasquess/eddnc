package eddn

import (
	"encoding/json"
	"testing"
	"time"
)

func TestValidate(t *testing.T) {
	u := Upload{Vaildate: true, TestUrl: true, DryRun: !testing.Verbose()}
	u.Http.Timeout = 6 * time.Second
	u.Header.Uploader = "_test_"
	u.Header.SwName = "goEDDNc"
	u.Header.SwVersion = "0.0.1"
	msg := make(map[string]interface{})
	err := json.Unmarshal([]byte(`{
    "systemName": "Munfayl",
    "stationName": "Samson",
    "timestamp": "2016-10-01T16:01:18Z",
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

func TestCommodityJ(t *testing.T) {
	u := Upload{Vaildate: true, TestUrl: true, DryRun: !testing.Verbose()}
	u.Http.Timeout = 6 * time.Second
	u.Header.Uploader = "_test_"
	u.Header.SwName = "goEDDNc"
	u.Header.SwVersion = "0.0.1"
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
