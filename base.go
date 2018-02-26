package document

import (
	"time"
)

const BaseType = "base"

type Base struct {
	Type             string    `json:"@type"`
	SubType          string    `json:"@subtype,omitempty"`
	Timestamp        time.Time `json:"@timestamp"`
	ID               string    `json:"@id,omitempty"`
	CertificateChain string    `json:"@certificateChain,omitempty"`
	Realm            string    `json:"@realm,omitempty"`
	raw              []byte
}

func NewBase() *Base {
	return &Base{
		Type:      BaseType,
		Timestamp: time.Now().UTC(),
	}
}

func (b *Base) GetTimestamp() time.Time {
	return b.Timestamp
}

func (b *Base) GetCertificateChain() string {
	return b.CertificateChain
}

func (b *Base) GetType() string {
	return b.Type
}

func (b *Base) GetSubType() string {
	return b.SubType
}

func (b *Base) GetRaw() []byte {
	return b.raw
}

func (b *Base) SetRaw(raw []byte) {
	b.raw = raw
}
