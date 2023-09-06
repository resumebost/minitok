package test

import (
	"minitok/cmd/user/tool"
	"testing"
)

func TestEncrypt(t *testing.T) {
	t.Log(tool.EncryptPassword("123456"))
}
