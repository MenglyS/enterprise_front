package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	Name      string
	Position  string
	Office    string
	Age       int
	StartDate string
	Salary    float64
}

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

	c.HTML(
		http.StatusOK,
		"dashboard.html",
		gin.H{
			"Title":  "Dashboard",
			"Style":  template.HTML(style),
			"Script": template.HTML(script),
		},
	)
}

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

	tableHeaders := []string{
		"Name",
		"Position",
		"Office",
		"Age",
		"Start Date",
		"Salary",
	}

	employees := make([]Employee, 50)

	for i := range employees {
		employees[i] = Employee{
			Name:      fmt.Sprintf("Employee %d", i+1),
			Position:  "Position",
			Office:    "Office",
			Age:       20 + i,
			StartDate: fmt.Sprintf("20%02d-01-01", i+1),
			Salary:    50000.00 + float64(i)*1000,
		}
	}

	c.HTML(
		http.StatusOK,
		"userManagement.html",
		gin.H{
			"Title":         "User Management",
			"TableName":     "Users List",
			"AddButton":     "Add User",
			"AddButtonIcon": "fa-user-plus",
			"TableHeaders":  tableHeaders,
			"TableData":     employees,
			"Style":         template.HTML(style),
			"Script":        template.HTML(script),
		},
	)
}

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
		"Name",
		"Position",
		"Office",
		"Age",
		"Start Date",
		"Salary",
	}

	employees := make([]Employee, 50)

	for i := range employees {
		employees[i] = Employee{
			Name:      fmt.Sprintf("Employee %d", i+1),
			Position:  "Position",
			Office:    "Office",
			Age:       20 + i,
			StartDate: fmt.Sprintf("20%02d-01-01", i+1),
			Salary:    50000.00 + float64(i)*1000,
		}
	}

	c.HTML(
		http.StatusOK,
		"job.html",
		gin.H{
			"Title":         "Job Management",
			"TableName":     "Jobs List",
			"AddButton":     "Add Job",
			"AddButtonIcon": "fa-plus",
			"TableHeaders":  tableHeaders,
			"TableData":     employees,
			"Style":         template.HTML(style),
			"Script":        template.HTML(script),
		},
	)
}

func LeaveManagement(c *gin.Context) {

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
		"Name",
		"Position",
		"Office",
		"Age",
		"Start Date",
		"Salary",
	}

	employees := make([]Employee, 50)

	for i := range employees {
		employees[i] = Employee{
			Name:      fmt.Sprintf("Employee %d", i+1),
			Position:  "Position",
			Office:    "Office",
			Age:       20 + i,
			StartDate: fmt.Sprintf("20%02d-01-01", i+1),
			Salary:    50000.00 + float64(i)*1000,
		}
	}

	c.HTML(
		http.StatusOK,
		"leave.html",
		gin.H{
			"Title":         "Leave Management",
			"TableName":     "Leaves List",
			"AddButton":     "Add Leave",
			"AddButtonIcon": "fa-plus",
			"TableHeaders":  tableHeaders,
			"TableData":     employees,
			"Style":         template.HTML(style),
			"Script":        template.HTML(script),
		},
	)
}

func ApplicantManagement(c *gin.Context) {

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
		"Name",
		"Position",
		"Office",
		"Age",
		"Start Date",
		"Salary",
	}

	employees := make([]Employee, 50)

	for i := range employees {
		employees[i] = Employee{
			Name:      fmt.Sprintf("Employee %d", i+1),
			Position:  "Position",
			Office:    "Office",
			Age:       20 + i,
			StartDate: fmt.Sprintf("20%02d-01-01", i+1),
			Salary:    50000.00 + float64(i)*1000,
		}
	}

	c.HTML(
		http.StatusOK,
		"applicant.html",
		gin.H{
			"Title":         "Applicant Management",
			"TableName":     "Applicants List",
			"AddButton":     "Add Applicant",
			"AddButtonIcon": "fa-plus",
			"TableHeaders":  tableHeaders,
			"TableData":     employees,
			"Style":         template.HTML(style),
			"Script":        template.HTML(script),
		},
	)
}

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
		"Name",
		"Position",
		"Office",
		"Age",
		"Start Date",
		"Salary",
	}

	employees := make([]Employee, 50)

	for i := range employees {
		employees[i] = Employee{
			Name:      fmt.Sprintf("Employee %d", i+1),
			Position:  "Position",
			Office:    "Office",
			Age:       20 + i,
			StartDate: fmt.Sprintf("20%02d-01-01", i+1),
			Salary:    50000.00 + float64(i)*1000,
		}
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
			"TableData":     employees,
			"Style":         template.HTML(style),
			"Script":        template.HTML(script),
		},
	)
}
