package main

import (
	"midterm/routers"

	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()

	route.Static("/app-assets", "./assets/app-assets")
	route.Static("/assets", "./assets/assets")
	route.Static("/gulp-tasks", "./assets/gulp-tasks")
	route.LoadHTMLGlob("views/**/*.html")

	routers.InitRoute(route)

	route.Run(":8080")
}
