package handler

import (
	"fmt"
	"minitok/cmd/api/rpc"
	"minitok/kitex_gen/favorite"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FavoriteAction(c *gin.Context) {

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
	ctx := c.Request.Context()
	req := &favorite.ActionRequest{
		VideoId:    reqData.VideoID,
		ActionType: reqData.ActionType,
	}

	resp, err := rpc.FavoriteAction(ctx, req)
	if err != nil {
		// 在 RPC 调用出错时记录错误信息
		fmt.Printf("Error in RPC call: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": resp.StatusCode,
			"status_msg":  resp.StatusMsg,
			"error":       "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": resp.StatusCode,
		"status_msg":  resp.StatusMsg,
	})

	fmt.Printf("请求参数: %+v\n", reqData)
	fmt.Printf("发起rpc请求: %+v\n", req)
	fmt.Printf("rpc响应: %+v\n", resp)
}

func FavoriteList(c *gin.Context) {

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
	ctx := c.Request.Context()
	req := &favorite.ListRequest{
		UserId: userID,
	}

	resp, err := rpc.FavoriteList(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": resp.StatusCode,
			"status_msg":  resp.StatusMsg,
			"error":       "Failed to fetch favorite list",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": resp.StatusCode,
		"status_msg":  resp.StatusMsg,
		"video_list":  resp.VideoList,
	})
}

