package handler

import (
	"github.com/gin-gonic/gin"
	"minitok/cmd/api/rpc"
	"minitok/internal/unierr"
	"minitok/kitex_gen/comment"
	"net/http"
	"strconv"
)

func CommentAction(c *gin.Context) {
	var req comment.ActionRequest

	req.Token = c.Query("token")
	var err error
	var actionTypeInt32 int64
	req.VideoId, err = strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionTypeInt32, err = strconv.ParseInt(c.Query("action_type"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": unierr.IllegalParams.ErrCode,
			"status_msg":  unierr.IllegalParams.ErrMsg,
		})
		return
	}
	req.ActionType = int32(actionTypeInt32)

	if req.ActionType == 1 {
		req.CommentText = c.Query("comment_text")
		if len(req.CommentText) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status_code": unierr.CommentEmpty.ErrCode,
				"status_msg":  unierr.CommentEmpty.ErrMsg,
			})
			return
		}
		//req.CommentId = -1
	} else if req.ActionType == 2 {
		req.CommentId, err = strconv.ParseInt(c.Query("comment_id"), 10, 64)
		if err != nil || req.CommentId <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status_code": unierr.IllegalParams.ErrCode,
				"status_msg":  unierr.IllegalParams.ErrMsg,
			})
			return
		}
		//req.CommentText = ""
	}

	resp, err := rpc.CommentAction(c, &req)
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
		"comment":     resp.Comment,
	})
}

func CommentList(c *gin.Context) {
	var req comment.ListRequest

	req.Token = c.Query("token")
	var err error
	req.VideoId, err = strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil || req.VideoId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": unierr.IllegalParams.ErrCode,
			"status_msg":  unierr.IllegalParams.ErrMsg,
		})
		return
	}

	resp, err := rpc.CommentList(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": resp.StatusCode,
			"status_msg":  resp.StatusMsg,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code":  resp.StatusCode,
		"status_msg":   resp.StatusMsg,
		"comment_list": resp.CommentList,
	})
}
