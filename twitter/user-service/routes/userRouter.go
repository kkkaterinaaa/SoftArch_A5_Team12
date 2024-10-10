package routes

import (
	"github.com/gin-gonic/gin"
	"twitter/user-service/controllers"
)

func SetupRouter(r *gin.Engine) {
	user := r.Group("users")
	user.GET("/:id", controllers.GetUser)
	user.POST("", controllers.CreateUser)
}
