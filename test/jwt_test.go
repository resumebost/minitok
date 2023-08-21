package test

import (
	"fmt"
	"minitok/internal/jwt"
	"testing"
)

func TestGenToken(t *testing.T) {
	str, _ := jwt.GenToken(1, "xiayi")
	fmt.Println(str)
}
