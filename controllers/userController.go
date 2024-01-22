package controllers

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"midterm/config"
	"midterm/models"
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

	resp, err := http.Get(config.Api + "user/getAll")
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

	c.HTML(
		http.StatusOK,
		"userManagement.html",
		gin.H{
			"Title":  "User Management",
			"Users":  result["data"],
			"Style":  template.HTML(style),
			"Script": template.HTML(script),
		},
	)
}
