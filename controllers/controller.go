package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {

	c.HTML(
		http.StatusOK,
		"dashboard.html",
		gin.H{},
	)
}
