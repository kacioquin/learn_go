package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kacioquin/learn_go/api_gin/api/controllers"
)

func AppRoutes(router *gin.Engine) *gin.RouterGroup {
	tweetController := controllers.NewTweetController()

	v1 := router.Group("/v1")
	{
		v1.GET("/tweets", tweetController.FindAll)
	}

	return v1
}
