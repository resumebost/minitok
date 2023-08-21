package main

import (
	"minitok/cmd/api/rpc"
	"minitok/internal/conf"
	"net/http"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
)

func initAll(e *gin.Engine) {
	initRouter(e)
	rpc.InitRPC()
}

func main() {
	r := gin.Default()

	initAll(r)

	if err := http.ListenAndServe(conf.APIServiceAddress, r); err != nil {
		klog.Fatal(err)
	}
}
