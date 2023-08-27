package handler

import (
	"github.com/gin-gonic/gin"
	"minitok/cmd/api/rpc"
	"minitok/cmd/user/model"
	"minitok/internal/unierr"
	"minitok/kitex_gen/user"
	"net/http"
	"strconv"
)

func UserRegister(c *gin.Context) {
	p := new(model.ParamUserRegister)
	//校验参数
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": unierr.IllegalParams.ErrCode,
			"status_msg":  unierr.IllegalParams.ErrMsg,
		})
		return
	}

	if len(p.Username) > 32 || len(p.Password) > 32 {
		c.JSON(http.StatusOK, gin.H{
			"status_code": unierr.UsernameOrPasswordLenMore32Characters.ErrCode,
			"status_msg":  unierr.UsernameOrPasswordLenMore32Characters.ErrMsg,
		})
		return
	}
	req := &user.RegisterRequest{
		Username: p.Username,
		Password: p.Password,
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
	p := new(model.ParamUserLogin)
	//校验参数
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": unierr.IllegalParams.ErrCode,
			"status_msg":  unierr.IllegalParams.ErrMsg,
		})
		return
	}
	if len(p.Username) > 32 || len(p.Password) > 32 {
		c.JSON(http.StatusOK, gin.H{
			"status_code": unierr.UsernameOrPasswordLenMore32Characters.ErrCode,
			"status_msg":  unierr.UsernameOrPasswordLenMore32Characters.ErrMsg,
		})
		return
	}

	req := &user.LoginRequest{
		Username: p.Username,
		Password: p.Password,
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
