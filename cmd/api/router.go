package main

import (
	"github.com/gin-gonic/gin"
	"minitok/cmd/api/handler"
	"minitok/internal/middleware"
)

func initRouter(r *gin.Engine) {
	router := r.Group("/douyin")

	router.POST("/user/register/", handler.UserRegister)
	router.POST("/user/login/", handler.UserLogin)
	router.GET("/feed/", handler.VideoFeed)

	// 后续接口由 AuthMiddleware 统一校验 jwt token
	router.Use(middleware.AuthMiddleware())

	router.GET("/user/", handler.UserInfo)

	router.POST("/publish/action/", handler.VideoPublishAction)
	router.GET("/publish/list/", handler.VideoPublishList)

	router.POST("/comment/action/", handler.CommentAction)
	router.GET("/comment/list/", handler.CommentList)

	router.POST("/favorite/action/", handler.FavoriteAction)
	router.GET("/favorite/list/", handler.FavoriteList)
}
