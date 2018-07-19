package eddn

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type Upload struct {
	vaildate bool
	header   struct {
		Uploader  string `json:"uploaderID"`
		SwName    string `json:"softwareName"`
		SwVersion string `json:"softwareVersion"`
	}
}

type eddnMsg struct {
	Schema  string      `json:"$schemaRef"`
	Header  interface{} `json:"header"`
	Message interface{} `json:"message"`
}

func (u *Upload) Journal(msg interface{}) error {
	emsg := eddnMsg{
		Schema:  ScmURIs[S_journal],
		Header:  &u.header,
		Message: msg,
	}
	jmsg, err := json.Marshal(&emsg)
	if err != nil {
		return err
	}
	if u.vaildate {
		errs := scmValidate(S_journal, jmsg)
		if len(errs) > 0 {
			buf := bytes.NewBuffer(nil)
			for _, e := range errs {
				fmt.Fprintln(buf, e.Error())
			}
			return errors.New(buf.String())
		}
	}
	return nil
}
