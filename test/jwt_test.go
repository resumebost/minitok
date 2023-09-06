package test

import (
	"fmt"
	"minitok/internal/jwt"
	"testing"
)

func TestGenToken(t *testing.T) {
	str, _ := jwt.GenToken(1, "xxhhy")
	fmt.Println(str)
}
