package tool

import (
	"crypto/sha256"
	"encoding/hex"
)

const secret = "minitok-secret"

func EncryptPassword(oPassword string) string {
	h := sha256.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
