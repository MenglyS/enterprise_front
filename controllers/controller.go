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
