package internal

import (
	"crypto/sha1"
	"encoding/hex"
)

// Hash returns the SHA1 hash of the input data as required by Mainline DHT.
func Hash(data []byte) []byte {
	digest := sha1.Sum(data)
	return digest[:]
}

// Hex returns the hex representation of the input data.
func Hex(data []byte) string {
	return hex.EncodeToString(data)
}
