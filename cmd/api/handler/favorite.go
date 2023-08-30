package handler

import (
	"fmt"
	"minitok/cmd/api/rpc"
	"minitok/internal/unierr"
	"minitok/kitex_gen/favorite"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func FavoriteAction(c *gin.Context) {

	token := c.GetHeader("Authorization")
	if len(token) > 7 && strings.ToLower(token[0:6]) == "bearer" {
		token = token[7:]
	}
	// Parse request data
	var reqData struct {
		VideoID    int64 `json:"video_id"`
		ActionType int32 `json:"action_type"`
	}
	if err := c.BindJSON(&reqData); err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	//RPC service
	req := &favorite.ActionRequest{
		Token:      token,
		VideoId:    reqData.VideoID,
		ActionType: reqData.ActionType,
	}

	resp, err := rpc.FavoriteAction(c, req)
	if err != nil {
		//在 RPC 调用出错时记录错误信息
		fmt.Printf("Error in RPC call: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": unierr.FavoriteActionError.ErrCode,
			"status_msg":  unierr.FavoriteActionError.ErrMsg,
			"error":       "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": resp.StatusCode,
		"status_msg":  resp.StatusMsg,
	})
}

func FavoriteList(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if len(token) > 7 && strings.ToLower(token[0:6]) == "bearer" {
		token = token[7:]
	}
	// Get UserID
	userIDStr := c.Query("user_id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	// Call RPC service
	req := &favorite.ListRequest{
		Token:  token,
		UserId: userID,
	}

	resp, err := rpc.FavoriteList(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code":  unierr.GetVideoListFiled.ErrCode,
			"status_msg":   unierr.GetVideoListFiled.ErrMsg,
			"error":       "Failed to fetch favorite list",
		})
		fmt.Println(resp)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": resp.StatusCode,
		"status_msg":  resp.StatusMsg,
		"video_list":  resp.VideoList,
	})
}
