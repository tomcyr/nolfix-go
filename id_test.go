// /home/tomcyr/code/nolfix/id_test.go
package nolfix_test

import (
	"testing"
	nolfix "github.com/tomcyr/nolfix-go"
)

func TestUUIDGeneratorProducesNonEmptyStrings(t *testing.T) {
	gen := nolfix.UUIDGenerator{}
	id1 := gen.NextID()
	id2 := gen.NextID()
	if id1 == "" {
		t.Error("expected non-empty ID")
	}
	if id1 == id2 {
		t.Errorf("expected unique IDs, got same: %s", id1)
	}
}
