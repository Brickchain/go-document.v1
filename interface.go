package document

import (
	"encoding/json"
	"time"
)

// Document describes the base types on a document
type Document interface {
	GetType() string
	GetSubType() string
	GetTimestamp() time.Time
	GetCertificateChain() string
	GetRaw() []byte
	SetRaw([]byte)
}

func Marshal(doc Document) []byte {
	docBytes, _ := json.Marshal(doc)
	return docBytes
}

func Unmarshal(data []byte) (Document, error) {
	base := &Base{}
	if err := json.Unmarshal(data, &base); err != nil {
		return base, err
	}

	base.SetRaw(data)

	var typ Document
	switch base.Type {
	case "base", "https://developer.brickchain.com/schemas/base.json":
		return base, nil
	case "action", "https://developer.brickchain.com/schemas/action.json":
		typ = &Action{}
	case "action-descriptor", "https://developer.brickchain.com/schemas/action-descriptor.json":
		typ = &ActionDescriptor{}
	case "certificate", "https://developer.brickchain.com/schemas/certificate.json":
		typ = &Certificate{}
	case "controller-binding", "https://developer.brickchain.com/schemas/controller-binding.json":
		typ = &ControllerBinding{}
	case "controller-descriptor", "https://developer.brickchain.com/schemas/controller-descriptor.json":
		typ = &ControllerDescriptor{}
	case "fact", "https://developer.brickchain.com/schemas/fact.json":
		typ = &Fact{}
	case "mandate", "https://developer.brickchain.com/schemas/mandate.json":
		typ = &Mandate{}
	case "mandate-token", "https://developer.brickchain.com/schemas/mandate-token.json":
		typ = &MandateToken{}
	case "message", "https://developer.brickchain.com/schemas/message.json":
		typ = &Message{}
	case "multipart", "https://developer.brickchain.com/schemas/multipart.json":
		typ = &Multipart{}
	case "realm-descriptor", "https://developer.brickchain.com/schemas/realm-descriptor.json":
		typ = &RealmDescriptor{}
	case "receipt", "https://developer.brickchain.com/schemas/receipt.json":
		typ = &Receipt{}
	case "revocation-checksum", "https://developer.brickchain.com/schemas/revocation-checksum.json":
		typ = &RevocationChecksum{}
	case "revocation-request", "https://developer.brickchain.com/schemas/revocation-request.json":
		typ = &RevocationRequest{}
	case "revocation", "https://developer.brickchain.com/schemas/revocation.json":
		typ = &Revocation{}
	case "scope-request", "https://developer.brickchain.com/schemas/scope-request.json":
		typ = &ScopeRequest{}
	case "signature-request", "https://developer.brickchain.com/schemas/signature-request.json":
		typ = &SignatureRequest{}
	default:
		return base, nil
	}

	if err := json.Unmarshal(data, &typ); err != nil {
		return base, err
	}
	typ.SetRaw(data)

	return typ, nil
}

func GetType(data []byte) (string, error) {
	b, err := Unmarshal(data)
	return b.GetType(), err
}

func GetSubType(data []byte) (string, error) {
	b, err := Unmarshal(data)
	return b.GetSubType(), err
}
