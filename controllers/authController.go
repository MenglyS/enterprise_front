package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"midterm/config"
	"net/http"
	"net/url"
	"strings"

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

func Login_Submit(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	data := url.Values{}
	data.Set("email", email)
	data.Set("password", password)

	req, err := http.NewRequest("POST", config.Api+"admin/login", strings.NewReader(data.Encode()))
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login?error="+url.QueryEscape(err.Error()))
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login?error="+url.QueryEscape(err.Error()))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.Redirect(http.StatusSeeOther, "/login?error=unauthorized")
		return
	}

	var response map[string]string
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login?error="+url.QueryEscape(err.Error()))
		return
	}

	fmt.Printf("Response: %v\n", response)

	c.SetCookie("token", response["token"], 3600, "/", "", false, true)

	c.Redirect(http.StatusSeeOther, "/")
}
