package handler

import (
	"github.com/gin-gonic/gin"
	"minitok/cmd/api/rpc"
	"minitok/internal/unierr"
	"minitok/internal/util"
	"minitok/kitex_gen/video"
	"net/http"
)

func VideoFeed(c *gin.Context) {
	var req video.FeedRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, err)
		return
	}

	//若时间为空则默认当前时间
	//TODO
	//解析时间

}

func VideoPublishAction(c *gin.Context) {
	var req video.PublishActionRequest

	req.Title = c.PostForm("title")
	req.Token = c.PostForm("token")

	fileHeader, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, unierr.VideoReadError)
		return
	}
	//将file转为[]byte
	req.Data, err = util.File2Bytes(fileHeader)
	if err != nil {
		c.JSON(http.StatusOK, unierr.VideoConvertError)
		return
	}

	if len(req.Title) == 0 || len(req.Data) == 0 {
		c.JSON(http.StatusOK, unierr.VideoIsEmptyError)
		return
	}

	resp, err := rpc.VideoPublishAction(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusOK, unierr.VideoPublishFiled)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func VideoPublishList(c *gin.Context) {
	var req video.PublishListRequest
	// BUG 应该用 c.Query 而不是 c.ShouldBind
	// GET 请求的参数通过 ?a=1&b=2 来传递
	// RPC 服务的连接在我本机测试没问题
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, err)
		return
	}

	if req.UserId <= 0 {
		c.JSON(http.StatusOK, unierr.IllegalParams)
		return
	}

	resp, err := rpc.VideoPublishList(c.Request.Context(), &video.PublishListRequest{
		UserId: req.UserId,
		Token:  req.Token,
	})
	if err != nil {
		c.JSON(http.StatusOK, unierr.GetPublishListFiled)
		return
	}
	c.JSON(http.StatusOK, resp)
}
