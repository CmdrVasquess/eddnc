package eddn

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const UploadURL = "https://eddn.edcd.io:4430/upload/"
const ConentType = "application/json; charset=utf-8"

func Ts(ts time.Time) string { return ts.Format(time.RFC3339) }

func NewMessage(ts string) map[string]interface{} {
	res := make(map[string]interface{})
	res["timestamp"] = ts
	return res
}

type Upload struct {
	Vaildate bool
	TestUrl  bool
	DryRun   bool
	Header   struct {
		Uploader  string `json:"uploaderID"`
		SwName    string `json:"softwareName"`
		SwVersion string `json:"softwareVersion"`
	}
	Http http.Client
}

type eddnMsg struct {
	Schema  string      `json:"$schemaRef"`
	Header  interface{} `json:"header"`
	Message interface{} `json:"message"`
}

func (u *Upload) Send(scm ScmId, msg interface{}) error {
	emsg := eddnMsg{
		Schema:  ScmURLs[scm],
		Header:  &u.Header,
		Message: msg,
	}
	if u.TestUrl {
		emsg.Schema = emsg.Schema + "/test"
	}
	jmsg, err := json.Marshal(&emsg)
	if err != nil {
		return err
	}
	if u.Vaildate {
		err := scmValidate(scm, jmsg)
		if err != nil {
			return err
		}
	}
	if !u.DryRun {
		rd := bytes.NewBuffer(jmsg)
		resp, err := u.Http.Post(UploadURL, ConentType, rd)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			msg, _ := ioutil.ReadAll(resp.Body)
			return fmt.Errorf("%d: %s", resp.StatusCode, string(msg))
		}
	}
	return nil
}
