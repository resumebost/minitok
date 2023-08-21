package test

import (
	"fmt"
	"minitok/internal/constant"
	"testing"
)

func TestConstant(t *testing.T) {
	constant.InitConstant()
	fmt.Println(constant.AllConstants.Datasource.DSNString())
}
