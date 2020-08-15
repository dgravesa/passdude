package useby

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

const saltLen = 32

func makeSalt() string {
	saltBytes := make([]byte, saltLen)
	rand.Read(saltBytes)
	return hex.EncodeToString(saltBytes)
}

func applySaltAndHash(in, salt string) string {
	saltedStr := in + "." + salt
	hashBytes := sha256.Sum256([]byte(saltedStr))
	return hex.EncodeToString(hashBytes[:])
}
