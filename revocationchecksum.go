package document

import (
	"time"
)

const RevocationChecksumType = "revocation-checksum"

type RevocationChecksum struct {
	Base
	Multihash string `json:"multihash"` // The signature from this document is to be revoked
}

func NewRevocationChecksum(multihash string) *RevocationChecksum {
	r := &RevocationChecksum{
		Base: Base{
			Type:      RevocationChecksumType,
			Timestamp: time.Now(),
		},
		Multihash: multihash,
	}
	return r
}
