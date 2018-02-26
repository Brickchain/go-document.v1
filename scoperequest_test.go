package document

import "testing"

func TestNewScopeRequest(t *testing.T) {
	s := NewScopeRequest(10)

	if s.Type != ScopeRequestType {
		t.Fatal("Wrong type: ", s.Type)
	}
}
