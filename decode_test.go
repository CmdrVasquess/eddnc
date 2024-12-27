package eddnc

import (
	"testing"
)

func TestScmMatch(t *testing.T) {
	scm := SchemaString([]byte(
		`{"$schemaRef": "https://eddn.edcd.io/schemas/journal/1", "header": {"gatewa`,
	))
	if scm != "https://eddn.edcd.io/schemas/journal/1" {
		t.Fatal("schema not found")
	}
}

func TestMessageEnd(t *testing.T) {
	raw := []byte(` {"foo": {"bar":"\"baz\""}["}}}}"]}`)
	me := MessageEnd(raw)
	msg := string(raw[:me])
	if msg != ` {"foo": {"bar":"\"baz\""}["}}}}"]}` {
		t.Fatalf("wrong message [%s]", msg)
	}
}
