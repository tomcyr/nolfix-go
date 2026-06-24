package nolfix_test

import (
	"testing"

	nolfix "github.com/tomcyr/nolfix-go"
)

func TestIntGeneratorProducesSequentialIDs(t *testing.T) {
	gen := &nolfix.IntGenerator{}
	id1 := gen.NextID()
	id2 := gen.NextID()
	if id1 != 1 {
		t.Errorf("expected first ID=1, got %d", id1)
	}
	if id2 != 2 {
		t.Errorf("expected second ID=2, got %d", id2)
	}
}
