package document

import (
	"testing"
)

func TestBase(t *testing.T) {
	b := NewBase()

	if b.Type != BaseType {
		t.Fatal("Wrong type")
	}
}
