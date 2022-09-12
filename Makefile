EDDNDIR:=./doc/EDDN
EDDNSCMS:=$(wildcard $(EDDNDIR)/schemas/*.json)
JSCMS:=$(patsubst %,./schemas/%,$(notdir $(EDDNSCMS)))

all: schemas.go

schemas.go: $(JSCMS)
	go run gen/main.go > $@
	gofmt -w $@
	go generate

schemas/%.json: $(EDDNDIR)/schemas/*.json
	cp $< $@
