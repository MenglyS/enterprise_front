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

func Dashboard(c *gin.Context) {

	const style = `
	<link rel="stylesheet" type="text/css" href="app-assets/css/vendors.css">
	<link rel="stylesheet" type="text/css" href="app-assets/vendors/css/charts/jquery-jvectormap-2.0.3.css">
	<link rel="stylesheet" type="text/css" href="app-assets/vendors/css/charts/morris.css">
	<link rel="stylesheet" type="text/css" href="app-assets/vendors/css/extensions/unslider.css">
	<link rel="stylesheet" type="text/css" href="app-assets/vendors/css/weather-icons/climacons.min.css">
	<link rel="stylesheet" type="text/css" href="app-assets/vendors/css/calendars/fullcalendar.min.css">
	<link rel="stylesheet" type="text/css" href="app-assets/css/core/menu/menu-types/vertical-overlay-menu.css">
	<link rel="stylesheet" type="text/css" href="app-assets/css/core/colors/palette-gradient.css">
	<link rel="stylesheet" type="text/css" href="app-assets/css/plugins/calendars/clndr.css">
	<link rel="stylesheet" type="text/css" href="app-assets/css/core/colors/palette-climacon.css">
	<link rel="stylesheet" type="text/css" href="app-assets/css/pages/users.css">
	<link rel="stylesheet" type="text/css" href="app-assets/css/plugins/calendars/fullcalendar.css">
	`

	const script = `
	<script src="app-assets/vendors/js/extensions/jquery.knob.min.js"></script>
	<script src="app-assets/vendors/js/charts/raphael-min.js"></script>
	<script src="app-assets/vendors/js/charts/morris.min.js"></script>
	<script src="app-assets/vendors/js/charts/chartist.min.js"></script>
	<script src="app-assets/vendors/js/charts/chartist-plugin-tooltip.js"></script>
	<script src="app-assets/vendors/js/charts/chart.min.js"></script>
	<script src="app-assets/vendors/js/charts/jquery.sparkline.min.js"></script>
	<script src="app-assets/vendors/js/extensions/moment.min.js"></script>
	<script src="app-assets/vendors/js/extensions/underscore-min.js"></script>
	<script src="app-assets/vendors/js/extensions/clndr.min.js"></script>
	<script src="app-assets/vendors/js/extensions/unslider-min.js"></script>
	<script src="app-assets/vendors/js/extensions/fullcalendar.min.js"></script>
	<script src="app-assets/js/scripts/pages/dashboard-project.js"></script>
	<script src="app-assets/js/scripts/extensions/fullcalendar.js"></script>
	`
	// Get all applicants
	resp, err := http.Get(config.Api + "applicant/getAll")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result map[string][]models.Applicant
	json.Unmarshal(body, &result)

	total := len(result["data"])

	res, err := http.Get(config.Api + "job/getAll")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var jobs map[string][]models.Job
	json.Unmarshal(body, &jobs)

	c.HTML(
		http.StatusOK,
		"dashboard.html",
		gin.H{
			"Title":           "Dashboard",
			"TotalApplicants": total,
			"Applicants":      result["data"],
			"Jobs":            jobs["data"],
			"Style":           template.HTML(style),
			"Script":          template.HTML(script),
		},
	)
}
