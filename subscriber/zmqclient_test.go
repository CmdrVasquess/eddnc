package subscriber

import (
	"testing"
)

func TestScmMatch(t *testing.T) {
	m := scmMatch.FindStringSubmatch(
		`{"$schemaRef": "https://eddn.edcd.io/schemas/journal/1", "header": {"gatewa`,
	)
	if m == nil {
		t.Fatal("no match")
	}
	if m[1] != "https://eddn.edcd.io/schemas/journal/1" {
		t.Errorf("wrong schema: '%s'", m[1])
	}
}
