package internal

import (
	"crypto/ed25519"
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

// GenerateKeypair generates a public/private keypair using ed25519.
func GenerateKeypair() (ed25519.PublicKey, ed25519.PrivateKey, error) {
	return ed25519.GenerateKey(nil)
}

// Sign returns the signature of the input data using the private key.
func Sign(privateKey ed25519.PrivateKey, data []byte) []byte {
	return ed25519.Sign(privateKey, data)
}

// Verify returns true if the signature is valid for the given data and public key.
func Verify(publicKey ed25519.PublicKey, data []byte, signature []byte) bool {
	return ed25519.Verify(publicKey, data, signature)
}
