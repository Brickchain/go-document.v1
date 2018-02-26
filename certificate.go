package document

import (
	"time"

	"strings"

	jose "gopkg.in/square/go-jose.v1"
)

const CertificateType = "certificate"

type Certificate struct {
	Base
	TTL           int              `json:"ttl,omitempty"`
	Issuer        *jose.JsonWebKey `json:"issuer,omitempty"`
	Subject       *jose.JsonWebKey `json:"subject,omitempty"`
	DocumentTypes []string         `json:"documentTypes,omitempty"`
	KeyLevel      int              `json:"keyLevel,omitempty"`
}

func NewCertificate(issuer, subject *jose.JsonWebKey, keyLevel int, ttl int) *Certificate {
	return &Certificate{
		Base: Base{
			Type:      CertificateType,
			Timestamp: time.Now().UTC(),
		},
		Issuer:        issuer,
		Subject:       subject,
		DocumentTypes: []string{"*"},
		TTL:           ttl,
		KeyLevel:      keyLevel,
	}
}

func (c *Certificate) HasExpired() bool {
	return time.Now().UTC().After(c.Timestamp.Add(time.Second * time.Duration(c.TTL)))
}

func (c *Certificate) AllowedType(doc Document) bool {
	for _, allowedType := range c.DocumentTypes {
		if strings.Contains(allowedType, "/") {
			parts := strings.Split(allowedType, "/")
			if len(parts) < 2 {
				return false
			}
			if parts[0] == "*" || parts[0] == doc.GetType() {
				if parts[1] == "*" || parts[1] == doc.GetSubType() {
					return true
				}
			}
		} else {
			if doc.GetType() == allowedType || allowedType == "*" {
				return true
			}
		}
	}
	return false
}
