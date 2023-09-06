package handler

import (
	"github.com/gin-gonic/gin"
	"minitok/cmd/api/rpc"
	"minitok/internal/unierr"
	"minitok/kitex_gen/user"
	"net/http"
	"strconv"
)

func UserRegister(c *gin.Context) {
	Username := c.Query("username")
	Password := c.Query("password")

	if len(Username) == 0 || len(Password) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status_code": unierr.UsernameOrPasswordIsEmpty.ErrCode,
			"status_msg":  unierr.UsernameOrPasswordIsEmpty.ErrMsg,
		})
		return
	}

	if len(Username) > 32 || len(Password) > 32 {
		c.JSON(http.StatusOK, gin.H{
			"status_code": unierr.UsernameOrPasswordLenMore32Characters.ErrCode,
			"status_msg":  unierr.UsernameOrPasswordLenMore32Characters.ErrMsg,
		})
		return
	}
	req := &user.RegisterRequest{
		Username: Username,
		Password: Password,
	}

	resp, err := rpc.Register(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusOK, unierr.InternalError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": resp.StatusCode,
		"status_msg":  resp.StatusMsg,
		"user_id":     resp.UserId,
		"token":       resp.Token,
	})
}

func UserLogin(c *gin.Context) {
	Username := c.Query("username")
	Password := c.Query("password")
	if len(Username) == 0 || len(Password) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status_code": unierr.UsernameOrPasswordIsEmpty.ErrCode,
			"status_msg":  unierr.UsernameOrPasswordIsEmpty.ErrMsg,
		})
		return
	}
	if len(Username) > 32 || len(Password) > 32 {
		c.JSON(http.StatusOK, gin.H{
			"status_code": unierr.UsernameOrPasswordLenMore32Characters.ErrCode,
			"status_msg":  unierr.UsernameOrPasswordLenMore32Characters.ErrMsg,
		})
		return
	}

	req := &user.LoginRequest{
		Username: Username,
		Password: Password,
	}

	resp, err := rpc.Login(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusOK, unierr.InternalError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": resp.StatusCode,
		"status_msg":  resp.StatusMsg,
		"user_id":     resp.UserId,
		"token":       resp.Token,
	})
}

func UserInfo(c *gin.Context) {
	userid := c.Query("user_id")
	token := c.Query("token")

	if token == "" {
		c.JSON(http.StatusOK, unierr.TokenNotExist)
	}

	userID, err := strconv.ParseInt(userid, 10, 64)

	if err != nil {
		c.JSON(http.StatusOK, unierr.UserIdInvalid)
	}

	req := &user.InfoRequest{
		UserId: userID,
		Token:  token,
	}

	resp, err := rpc.Info(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusOK, unierr.InternalError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": resp.StatusCode,
		"status_msg":  resp.StatusMsg,
		"user":        resp.User,
	})
}
