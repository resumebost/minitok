package handler

import (
	"github.com/gin-gonic/gin"
	"minitok/cmd/api/rpc"
	"minitok/internal/unierr"
	"minitok/internal/util"
	"minitok/kitex_gen/video"
	"net/http"
	"strconv"
	"time"
)

func VideoFeed(c *gin.Context) {
	var req video.FeedRequest

	req.Token = c.Query("token")
	timeStr := c.Query("latest_time")

	//若时间为空则默认当前时间
	if len(timeStr) == 0 {
		req.LatestTime = time.Now().UTC().UnixMilli() //实际为毫秒
	} else {
		var err error
		req.LatestTime, err = strconv.ParseInt(timeStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status_code": unierr.IllegalParams.ErrCode,
				"status_msg":  unierr.IllegalParams.ErrMsg,
			})
			return
		}
	}

	resp, err := rpc.Feed(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": resp.StatusCode,
			"status_msg":  resp.StatusMsg,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": resp.StatusCode,
		"status_msg":  resp.StatusMsg,
		"video_list":  resp.VideoList,
		"next_time":   resp.NextTime,
	})
}

func VideoPublishAction(c *gin.Context) {
	var req video.PublishActionRequest

	req.Title = c.PostForm("title")
	req.Token = c.PostForm("token")

	fileHeader, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": unierr.IllegalParams.ErrCode,
			"status_msg":  unierr.IllegalParams.ErrMsg,
		})
		return
	}
	//将file转为[]byte
	req.Data, err = util.File2Bytes(fileHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": unierr.VideoConvertError.ErrCode,
			"status_msg":  unierr.VideoConvertError.ErrMsg,
		})
		return
	}

	if len(req.Title) == 0 || len(req.Data) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": unierr.IllegalParams.ErrCode,
			"status_msg":  unierr.IllegalParams.ErrMsg,
		})
		return
	}

	resp, err := rpc.VideoPublishAction(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": resp.StatusCode,
			"status_msg":  resp.StatusMsg,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": resp.StatusCode,
		"status_msg":  resp.StatusMsg,
	})
}

func VideoPublishList(c *gin.Context) {
	var req video.PublishListRequest

	idStr := c.Query("user_id")
	req.Token = c.Query("token")

	var err error
	req.UserId, err = strconv.ParseInt(idStr, 10, 64)

	if err != nil || req.UserId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": unierr.IllegalParams.ErrCode,
			"status_msg":  unierr.IllegalParams.ErrMsg,
		})
		return
	}

	resp, err := rpc.VideoPublishList(c, &video.PublishListRequest{
		UserId: req.UserId,
		Token:  req.Token,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": unierr.GetPublishListFiled.ErrCode,
			"status_msg":  unierr.GetPublishListFiled.ErrMsg,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": resp.StatusCode,
		"status_msg":  resp.StatusMsg,
		"video_list":  resp.VideoList,
	})
}
