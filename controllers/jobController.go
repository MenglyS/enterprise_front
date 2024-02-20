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

func JobManagement(c *gin.Context) {

	const style = `
	<link rel="stylesheet" type="text/css" href="app-assets/css/vendors.css">
    <link rel="stylesheet" type="text/css" href="app-assets/vendors/css/tables/datatable/datatables.min.css">
	<link rel="stylesheet" type="text/css" href="app-assets/css/core/menu/menu-types/vertical-overlay-menu.css">
    <link rel="stylesheet" type="text/css" href="app-assets/css/core/colors/palette-gradient.css">
	`

	const script = `
	<script src="app-assets/vendors/js/tables/datatable/datatables.min.js"></script>
	<script src="app-assets/js/scripts/tables/datatables/datatable-basic.js"></script>
	`

	tableHeaders := []string{
		"Title",
		"Category",
		"Announcement Image",
		"Status",
		"Action",
	}
	token, err := c.Cookie("token")
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login?error="+err.Error())
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", config.Api+"job/getAll", nil)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login?error="+err.Error())
		return
	}

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login?error="+err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result map[string][]models.Job
	json.Unmarshal(body, &result)

	user, err := GetAuthUser(token)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login?error="+err.Error())
		return

	}

	c.HTML(
		http.StatusOK,
		"job.html",
		gin.H{
			"Title":         "Job Management",
			"TableName":     "Jobs List",
			"AddButton":     "Add Job",
			"AddButtonIcon": "fa-plus",
			"User":          user,
			"TableHeaders":  tableHeaders,
			"TableData":     result["data"],
			"Style":         template.HTML(style),
			"Script":        template.HTML(script),
		},
	)
}

func Create_Job(c *gin.Context) {

	token, err := c.Cookie("token")
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login?error="+err.Error())
		return
	}
	title := c.PostForm("title")
	cat := c.PostForm("cat")
	des := c.PostForm("des")
	contact := c.PostForm("contact")
	exp := c.PostForm("exp-date")

	// Create a buffer
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)

	// Add the other form data
	_ = writer.WriteField("Title", title)
	_ = writer.WriteField("CategoryIds", cat)
	_ = writer.WriteField("Description", des)
	_ = writer.WriteField("Contact", contact)
	_ = writer.WriteField("ExpiryDate", exp)

	// Open the file

	file, err := c.FormFile("img")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	src, err := file.Open()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer src.Close()

	// Add the file to the form
	part, err := writer.CreateFormFile("image", file.Filename)
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

	// Now you can send the request with the buffer as the body
	req, err := http.NewRequest("POST", config.Api+"job/create/", buf)
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

	c.Redirect(http.StatusSeeOther, "/jobs")
}
