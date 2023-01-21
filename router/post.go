package router

import (
	"gin/basic/controller"

	"github.com/gin-gonic/gin"
)

func PostRouter(engine *gin.Engine) {
	engine.GET("/post", controller.GetPost) //getPost this will come from controller
	engine.POST("/post", controller.CreatePost)

	//engine.PUT("/post", controller.GetPost)
	engine.DELETE("/post", controller.DeletePost)
	//engine.PATCH("/post", controller.GetPost)
}
