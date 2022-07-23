package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	ginx "family-joint-school/common/http/gin"
	middleware "family-joint-school/common/http/gin/middware"
	"family-joint-school/handler/http/api"
)

// Router HTTP 路由
func Router(apiHandler api.Handler) *gin.Engine {
	engine := gin.Default()
	engine.Use(middleware.Recover(), middleware.NewLogger(), middleware.AccessHandler())
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	engine.POST("/upload", apiHandler.Upload)
	engine.POST("/login", ginx.HandlerFuncWrapper(apiHandler.Login))
	{
		apiGroup := engine.Group("/api", middleware.JwtAuth())
		apiGroup.POST("/class/list", ginx.HandlerFuncWrapper(apiHandler.ClassList))
		apiGroup.POST("/notice/list", ginx.HandlerFuncWrapper(apiHandler.NoticeList))
		apiGroup.POST("/homework/list", ginx.HandlerFuncWrapper(apiHandler.HomeWorkList))
		apiGroup.POST("/homework/submit", ginx.HandlerFuncWrapper(apiHandler.HomeWorkSubmit))
		apiGroup.POST("/homework/user/detail", ginx.HandlerFuncWrapper(apiHandler.UserHomeWorkDetail))

	}
	{
		cmsGroup := engine.Group("/cms", middleware.JwtAuthCMS())
		cmsGroup.POST("/notice/publish", ginx.HandlerFuncWrapper(apiHandler.NoticePublish))
		cmsGroup.POST("/homework/publish", ginx.HandlerFuncWrapper(apiHandler.HomeWorkPublish))
		cmsGroup.POST("/homework/user/list", ginx.HandlerFuncWrapper(apiHandler.UserHomeWorkList))
	}
	return engine
}
