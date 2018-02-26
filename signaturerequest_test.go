package document

import "testing"

func TestNewSignatureRequest(t *testing.T) {
	s := NewSignatureRequest(10)

	if s.Type != SignatureRequestType {
		t.Fatal("Wrong type: ", s.Type)
	}
}
