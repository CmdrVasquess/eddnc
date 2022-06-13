EDDNDIR:=./doc/EDDN
JSCMS:=$(wildcard $(EDDNDIR)/schemas/*.json)

all: schemas.go

schemas.go: $(JSCMS)
	go run gen/main.go > $@
	gofmt -w $@
	go generate

#JSCMGO:=$(patsubst %.json,schema/%.go,$(notdir $(JSCMS)))
#
#static-schemas: $(JSCMGO)
#
#schema/%.go: $(EDDNDIR)/schemas/%.json
#	gojsonschema -p schema -o $@ $<
