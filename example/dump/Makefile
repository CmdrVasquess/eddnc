msgs:=$(shell find msgs -name '*.json')
jsums:=$(patsubst %.json,%.jsum,$(msgs))

.PHONY: clean

all: $(jsums)

clean:
	rm -f $(jsums)

%.jsum: %.json
	jsum $< > $@
