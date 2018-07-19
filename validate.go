package eddn

import (
	"bytes"
	"errors"
	"fmt"

	//jscm "github.com/qri-io/jsonschema"
	jscm "github.com/xeipuuv/gojsonschema"
)

//go:generate ./genschemas.sh
//go:generate stringer -type ScmId

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

type ScmId int

func scmValidate(scmId ScmId, json []byte) error {
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
