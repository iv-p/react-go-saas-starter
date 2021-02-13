package id

import (
	"github.com/rs/xid"
)

// ID is a unique identifier
type ID = string

// New returns an unique identifier
func New() ID {
	return xid.New().String()
}
