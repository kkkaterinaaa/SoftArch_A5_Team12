package routes

import (
	"github.com/gin-gonic/gin"
	"twitter/interction-service/controllers"
)

func SetupRouter(r *gin.Engine) {
	user := r.Group("likes")
	user.POST("", controllers.LikeMessage)
	user.GET("/message/:messageID", controllers.GetMessageLikes)
}
