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
