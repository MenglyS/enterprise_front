package main

import (
	"midterm/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	route.LoadHTMLGlob("views/**/*.html")

	routers.InitRoute(route)

	route.Run(":8080")
}
