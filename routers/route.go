package routers

import (
	"midterm/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoute(route *gin.Engine) {
	route.GET("/", controllers.Dashboard)
	route.GET("/users", controllers.UserManagement)
	route.POST("/user/create", controllers.Create_User)
	route.GET("/jobs", controllers.JobManagement)
	route.POST("/job/create", controllers.Create_Job)
	route.GET("/expenses", controllers.ExpendManagement)
	route.POST("/expenses/edit/:id", controllers.Edit_Expense)
	route.GET("/leaves", controllers.LeaveManagement)
	route.POST("/leave/edit/:id", controllers.Edit_Leave)
	route.GET("/applicants", controllers.ApplicantManagement)
	route.GET("/login", controllers.Login)
	route.POST("/login/submit", controllers.Login_Submit)
	route.POST("/applicant/edit/:id", controllers.Edit_Applicant)
}
