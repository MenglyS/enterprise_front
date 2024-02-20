package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"midterm/config"
	"midterm/models"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserManagement(c *gin.Context) {

	const style = `
    <link rel="stylesheet" type="text/css" href="app-assets/css/vendors.css">
    <link rel="stylesheet" type="text/css" href="app-assets/vendors/css/tables/datatable/dataTables.bootstrap4.min.css">
    <link rel="stylesheet" type="text/css" href="app-assets/vendors/css/tables/extensions/rowReorder.dataTables.min.css">
    <link rel="stylesheet" type="text/css" href="app-assets/vendors/css/tables/extensions/responsive.dataTables.min.css">
    <link rel="stylesheet" type="text/css" href="app-assets/vendors/css/forms/icheck/icheck.css">
    <link rel="stylesheet" type="text/css" href="app-assets/vendors/css/forms/icheck/custom.css">
	<link rel="stylesheet" type="text/css" href="app-assets/css/core/menu/menu-types/vertical-overlay-menu.css">
    <link rel="stylesheet" type="text/css" href="app-assets/css/core/colors/palette-gradient.css">
    <link rel="stylesheet" type="text/css" href="app-assets/css/pages/users.css">
	`

	const script = `
    <script src="app-assets/vendors/js/tables/jquery.dataTables.min.js"></script>
    <script src="app-assets/vendors/js/tables/datatable/dataTables.bootstrap4.min.js"></script>
    <script src="app-assets/vendors/js/tables/datatable/dataTables.responsive.min.js"></script>
    <script src="app-assets/vendors/js/tables/datatable/dataTables.rowReorder.min.js"></script>
    <script src="app-assets/vendors/js/forms/icheck/icheck.min.js"></script>
	<script src="app-assets/js/scripts/pages/users-contacts.js"></script>
	`
	token, err := c.Cookie("token")
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login?error="+err.Error())
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", config.Api+"user/getAll", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Set the authorization header
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result map[string][]models.User
	json.Unmarshal(body, &result)

	user, err := GetAuthUser(token)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login?error="+err.Error())
		return

	}

	c.HTML(
		http.StatusOK,
		"userManagement.html",
		gin.H{
			"Title":  "User Management",
			"Users":  result["data"],
			"User":   user,
			"Style":  template.HTML(style),
			"Script": template.HTML(script),
		},
	)
}

func Create_User(c *gin.Context) {

	token, err := c.Cookie("token")
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login?error="+err.Error())
		return
	}
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
	role := c.PostForm("role")
	position := c.PostForm("position")
	salary := c.PostForm("salary")
	dob := c.PostForm("dob")
	phone := c.PostForm("phone")

	// Create a buffer
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)

	// Add the other form data
	_ = writer.WriteField("Username", username)
	_ = writer.WriteField("Email", email)
	_ = writer.WriteField("Password", password)
	_ = writer.WriteField("RoleId", role)
	_ = writer.WriteField("PositionId", position)
	_ = writer.WriteField("Salary", salary)
	_ = writer.WriteField("Dob", dob)
	_ = writer.WriteField("Phone", phone)

	// Open the file

	file, err := c.FormFile("img")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		src, err := file.Open()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
		defer src.Close()

		// Add the file to the form
		part, err := writer.CreateFormFile("profile", file.Filename)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
		_, err = io.Copy(part, src)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}

		// Close the multipart writer
		err = writer.Close()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
	}

	// Now you can send the request with the buffer as the body
	req, err := http.NewRequest("POST", config.Api+"user/create/", buf)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	// Don't forget to set the Content-Type header with the boundary
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+token)

	// Send the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer res.Body.Close()

	c.Redirect(http.StatusSeeOther, "/users")
}

func GetAuthUser(token string) (models.User, error) {
	var user models.User
	client := &http.Client{}
	req, err := http.NewRequest("GET", config.Api+"user/getProfile", nil)
	if err != nil {
		return user, err
	}

	// Set the authorization header
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		return user, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result map[string][]models.User
	json.Unmarshal(body, &result)

	fmt.Printf("result: %v\n", result)

	user = result["data"][0]

	return user, nil
}
