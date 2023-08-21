package main

import (
	"minitok/cmd/api/rpc"
	"minitok/internal/constant"
	"net/http"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
)

var apiConstants *constant.ServiceInfo

func initAll(e *gin.Engine) {
	constant.InitConstant()
	apiConstants = &constant.AllConstants.APIService

	initRouter(e)
	rpc.InitRPC()
}

func main() {
	r := gin.Default()

	initAll(r)

	if err := http.ListenAndServe(apiConstants.Addr(), r); err != nil {
		klog.Fatalf("Gin HTTP server failed to start: %v", err)
	}
}
