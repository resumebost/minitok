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
	return
}

func VideoPublishList(c *gin.Context) {
}
