package tool

import (
	"crypto/md5"
	"encoding/hex"
)

const secret = "minitok-secret"

func EncryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
