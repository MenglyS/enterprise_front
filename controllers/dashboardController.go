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
	<script src="app-assets/vendors/js/charts/chartist-plugin-tooltip.js"></script>
	<script src="app-assets/vendors/js/charts/chart.min.js"></script>
	<script src="app-assets/vendors/js/charts/jquery.sparkline.min.js"></script>
	<script src="app-assets/vendors/js/extensions/moment.min.js"></script>
	<script src="app-assets/vendors/js/extensions/clndr.min.js"></script>
	<script src="app-assets/vendors/js/extensions/unslider-min.js"></script>
	<script src="app-assets/vendors/js/extensions/fullcalendar.min.js"></script>
	<script src="app-assets/js/scripts/pages/dashboard-project.js"></script>
	<script src="app-assets/js/scripts/extensions/fullcalendar.js"></script>
	`

	token, err := c.Cookie("token")
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login?error="+err.Error())
		return
	}
	// Get all applicants
	client := &http.Client{}
	req, err := http.NewRequest("GET", config.Api+"applicant/getAll", nil)
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

	var result map[string][]models.Applicant
	json.Unmarshal(body, &result)

	total := len(result["data"])

	client = &http.Client{}
	req, err = http.NewRequest("GET", config.Api+"job/getAll", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Set the authorization header
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var jobs map[string][]models.Job
	json.Unmarshal(body, &jobs)

	client = &http.Client{}
	req, err = http.NewRequest("GET", config.Api+"expense/pendingCount", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Set the authorization header
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var expense map[string]int32
	json.Unmarshal(body, &expense)

	client = &http.Client{}
	req, err = http.NewRequest("GET", config.Api+"leave/pendingCount", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Set the authorization header
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var leave map[string]int32
	json.Unmarshal(body, &leave)

	client = &http.Client{}
	req, err = http.NewRequest("GET", config.Api+"applicant/scheduledCount", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Set the authorization header
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var scheduled map[string]int32
	json.Unmarshal(body, &scheduled)

	user, err := GetAuthUser(token)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login?error="+err.Error())
		return

	}

	c.HTML(
		http.StatusOK,
		"dashboard.html",
		gin.H{
			"Title":              "Dashboard",
			"TotalApplicants":    total,
			"Applicants":         result["data"],
			"Jobs":               jobs["data"],
			"Expenses":           expense["data"],
			"Leaves":             leave["data"],
			"User":               user,
			"ScheduledApplicant": scheduled["data"],
			"Style":              template.HTML(style),
			"Script":             template.HTML(script),
		},
	)
}
