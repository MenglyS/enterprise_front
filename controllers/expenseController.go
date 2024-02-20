package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"midterm/config"
	"midterm/models"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func ExpendManagement(c *gin.Context) {

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
		"Amount",
		"File",
		"Status",
		"Action",
	}
	token, err := c.Cookie("token")
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login?error="+err.Error())
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", config.Api+"expense/getAll", nil)
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

	var result map[string][]models.Expense
	json.Unmarshal(body, &result)

	user, err := GetAuthUser(token)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login?error="+err.Error())
		return

	}

	c.HTML(
		http.StatusOK,
		"expend.html",
		gin.H{
			"Title":         "Expenses Management",
			"TableName":     "Expend Request List",
			"AddButton":     "Add",
			"AddButtonIcon": "fa-plus",
			"TableHeaders":  tableHeaders,
			"User":          user,
			"TableData":     result["data"],
			"Style":         template.HTML(style),
			"Script":        template.HTML(script),
		},
	)
}

func Edit_Expense(c *gin.Context) {
	status := c.PostForm("Status")
	token, err := c.Cookie("token")
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login?error="+err.Error())
		return
	}

	data := url.Values{}
	data.Set("Status", status)

	id := c.Param("id")

	fmt.Printf("id: %s\n", id)

	req, err := http.NewRequest("PUT", config.Api+"expense/edit/"+id, strings.NewReader(data.Encode()))
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login?error="+url.QueryEscape(err.Error()))
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+token)

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

	c.Redirect(http.StatusSeeOther, "/expenses?success=record has been editted")

}
