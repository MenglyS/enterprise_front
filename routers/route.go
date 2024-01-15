package routers

import (
	"midterm/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoute(route *gin.Engine) {
	route.GET("/", controllers.Dashboard)
}
