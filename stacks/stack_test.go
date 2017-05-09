package stacks

import (
	"testing"
)

func TestStackIsEmpty(t *testing.T) {
	s := NewStack()

	if !s.IsEmpty() {
		t.Errorf("Stack empty: %t", s.IsEmpty())
	}
}
