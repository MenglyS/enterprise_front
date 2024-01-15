package controllers

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	const style = `
	<link rel="stylesheet" type="text/css" href="app-assets/css/vendors.css">
    <link rel="stylesheet" type="text/css" href="app-assets/vendors/css/forms/icheck/icheck.css">
    <link rel="stylesheet" type="text/css" href="app-assets/vendors/css/forms/icheck/custom.css">
	<link rel="stylesheet" type="text/css" href="app-assets/css/core/menu/menu-types/vertical-overlay-menu.css">
    <link rel="stylesheet" type="text/css" href="app-assets/css/core/colors/palette-gradient.css">
    <link rel="stylesheet" type="text/css" href="app-assets/css/pages/login-register.css">
	`

	const script = `
	<script src="app-assets/vendors/js/forms/validation/jqBootstrapValidation.js"></script>
    <script src="app-assets/vendors/js/forms/icheck/icheck.min.js"></script>
	<script src="app-assets/js/scripts/forms/form-login-register.js"></script>
	`
	c.HTML(
		http.StatusOK,
		"login.html",
		gin.H{
			"Title":  "Login",
			"Style":  template.HTML(style),
			"Script": template.HTML(script),
		},
	)
}
