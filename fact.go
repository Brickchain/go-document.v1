package document

import (
	"time"

	jose "gopkg.in/square/go-jose.v1"

	"github.com/spf13/cast"
)

const FactType = "fact"

type Fact struct {
	Base
	TTL       time.Duration          `json:"ttl,omitempty"`
	Issuer    string                 `json:"iss,omitempty"`
	Label     string                 `json:"label,omitempty"`
	Data      map[string]interface{} `json:"data"`
	Recipient *jose.JsonWebKey       `json:"recipient,omitempty"`
}

func NewFact(subType string) *Fact {
	return &Fact{
		Base: Base{
			Type:      FactType,
			Timestamp: time.Now(),
			SubType:   subType,
		},
		Data: make(map[string]interface{}),
	}
}

func (f *Fact) Set(key string, value interface{}) bool {
	_, prs := f.Data[key]
	f.Data[key] = value
	return !prs
}

func (f *Fact) Get(key string) interface{} {
	i, _ := f.Data[key]
	return i
}

func (f *Fact) GetString(key string) string {
	return cast.ToString(f.Get(key))
}
