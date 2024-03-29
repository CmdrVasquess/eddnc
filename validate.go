package eddnc

import (
	"bytes"
	"errors"
	"fmt"

	//jscm "github.com/qri-io/jsonschema"
	jscm "github.com/xeipuuv/gojsonschema"
)

func init() {
	for i := range ScmDefs {
		ld := jscm.NewStringLoader(ScmDefs[i])
		scm, err := jscm.NewSchema(ld)
		if err != nil {
			panic(err)
		}
		scmLs = append(scmLs, scm)
	}
}

var scmLs []*jscm.Schema

type ScmID int

func scmValidate(scmId ScmID, json []byte) error {
	scm := scmLs[scmId]
	ld := jscm.NewBytesLoader(json)
	res, err := scm.Validate(ld)
	if err != nil {
		panic(err)
	}
	if !res.Valid() {
		buf := bytes.NewBuffer(nil)
		for _, jerr := range res.Errors() {
			fmt.Fprintf(buf, "%s:%s:%s\n",
				scmId,
				jerr.Context().String(),
				jerr.String())
		}
		return errors.New(buf.String())
	}
	return nil
}
