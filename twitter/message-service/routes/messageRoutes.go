package routes

import (
	"github.com/gin-gonic/gin"
	"twitter/message-service/controllers"
)

func SetupRouter(r *gin.Engine) {
	user := r.Group("messages")
	user.GET("", controllers.GetFeed)
	user.POST("", controllers.CreateMessage)
}
