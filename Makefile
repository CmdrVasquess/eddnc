EDDNDIR:=./doc/EDDN
JSCMS:=$(wildcard $(EDDNDIR)/schemas/*.json)

all: schema/schemas.go

schema/schemas.go:
	./genschemas.sh

#JSCMGO:=$(patsubst %.json,schema/%.go,$(notdir $(JSCMS)))
#
#static-schemas: $(JSCMGO)
#
#schema/%.go: $(EDDNDIR)/schemas/%.json
#	gojsonschema -p schema -o $@ $<
