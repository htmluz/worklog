package idgen

import (
	"crypto/rand"
	"encoding/hex"
)

func New() string {
	bytes := make([]byte, 6)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
