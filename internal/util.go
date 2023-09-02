package internal

import (
	"crypto/sha1"
	"encoding/hex"

	"github.com/tv42/zbase32"
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

// Z32Encode returns the zbase32 representation of the input data.
func Z32Encode(data []byte) string {
	return zbase32.EncodeToString(data)
}

// Z32Decode returns the decoded zbase32 representation of the input data.
func Z32Decode(data string) ([]byte, error) {
	return zbase32.DecodeString(data)
}
