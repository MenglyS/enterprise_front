package routers

import (
	"midterm/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoute(route *gin.Engine) {
	route.GET("/", controllers.Dashboard)
	route.GET("/users", controllers.UserManagement)
	route.GET("/jobs", controllers.JobManagement)
	route.GET("/expenses", controllers.ExpendManagement)
	route.GET("/leaves", controllers.LeaveManagement)
	route.GET("/applicants", controllers.ApplicantManagement)
	route.GET("/login", controllers.Login)
}
