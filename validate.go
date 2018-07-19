package eddn

import (
	jscm "github.com/qri-io/jsonschema"
)

//go:generate ./genschemas.sh

func init() {
	for i := range ScmDefs {
		s := &jscm.Schema{}
		err := s.UnmarshalJSON([]byte(ScmDefs[i]))
		if err != nil {
			panic(err)
		}
		//		scmRoots = append(scmRoots, jscm.Must(ScmDefs[i]))
	}
}

var scmRoots []*jscm.RootSchema

type ScmId int

func scmValidate(scmId ScmId, json []byte) []jscm.ValError {
	scm := scmRoots[scmId]
	res, err := scm.ValidateBytes(json)
	if err != nil {
		panic(err)
	}
	return res
}
