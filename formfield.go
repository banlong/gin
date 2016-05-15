package main
import (
	"net/http"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := registerRoutes()
	r.Run(":3000")
}


func registerRoutes() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.GET("/employees/:id/vacation", func(c *gin.Context) {
		id := c.Param("id")
		timesOff, ok := TimesOff[id]

		if !ok {
			c.String(http.StatusNotFound, "404 - Page Not Found")
			return
		}

		c.HTML(http.StatusOK, "vacation-overview.html",
			map[string]interface{}{
				"TimesOff": timesOff,
			})
	})

	admin := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	}))
	admin.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin-overview.html",
			map[string]interface{}{
				"Employees": employees,
			})
	})

	admin.GET("/employees/:id", func(c *gin.Context) {
		id := c.Param("id")
		if id == "add" {
			c.HTML(http.StatusOK, "admin-employee-add.html", nil)
			return
		}
		employee, ok := employees[id]

		if !ok {
			c.String(http.StatusNotFound, "404 - Not Found")
			return
		}

		c.HTML(http.StatusOK, "admin-employee-edit.html",
			map[string]interface{}{
				"Employee": employee,
			})
	})

	admin.POST("/employees/:id", func(c *gin.Context) {
		id := c.Param("id") //this is the action for example "add" it to add user
		log.Println("id == add: ", id == "add")
		if id == "add" {
			pto, err := strconv.ParseFloat(c.PostForm("pto"), 32)
			if err != nil {
				c.String(http.StatusBadRequest, err.Error())
				return
			}

			startDate, err := time.Parse("2006-01-02", c.PostForm("startDate"))
			if err != nil {
				c.String(http.StatusBadRequest, err.Error())
				return
			}

			var emp Employee
			emp.ID = 42
			emp.FirstName = c.PostForm("firstName")
			emp.LastName = c.PostForm("lastName")
			emp.Position = c.PostForm("position")
			emp.Status = "Active"
			emp.TotalPTO = float32(pto)
			emp.StartDate = startDate
			employees["42"] = emp

			log.Println("Employee", emp)

			c.Redirect(http.StatusMovedPermanently, "/admin/employees/42")
		}
	})

	r.Static("/public", "./public")

	return r
}


var employees = map[string]Employee{
	"962134": Employee{
		ID:        962134,
		FirstName: "Jennifer",
		LastName:  "Watson",
		Position:  "CEO",
		StartDate: time.Now().Add(-13 * time.Hour * 24 * 365),
		Status:    "Active",
		TotalPTO:  30,
	},
	"176158": Employee{
		ID:        176158,
		FirstName: "Allison",
		LastName:  "Jane",
		Position:  "COO",
		StartDate: time.Now().Add(-4 * time.Hour * 24 * 365),
		Status:    "Active",
		TotalPTO:  20,
	},
	"160898": Employee{
		ID:        160898,
		FirstName: "Aakar",
		LastName:  "Uppal",
		Position:  "CTO",
		StartDate: time.Now().Add(-6 * time.Hour * 24 * 365),
		TotalPTO:  20,
	},
	"297365": Employee{
		ID:        297365,
		FirstName: "Jonathon",
		LastName:  "Anderson",
		Position:  "Worker Bee",
		StartDate: time.Now().Add(-12 * time.Hour * 24 * 365),
		TotalPTO:  30,
	},
}

var TimesOff = map[string][]TimeOff{
	"962134": []TimeOff{
		{
			Type:      "Holiday",
			Amount:    8.,
			StartDate: time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC),
			Status:    "Taken",
		}, {
			Type:      "PTO",
			Amount:    16.,
			StartDate: time.Date(2016, 8, 16, 0, 0, 0, 0, time.UTC),
			Status:    "Scheduled",
		}, {
			Type:      "PTO",
			Amount:    16.,
			StartDate: time.Date(2016, 12, 8, 0, 0, 0, 0, time.UTC),
			Status:    "Requested",
		},
	},
}

type Employee struct {
	ID        uint
	FirstName string
	LastName  string
	StartDate time.Time
	Position  string
	TotalPTO  float32
	Status    string
	TimesOff  []TimeOff
}

type TimeOff struct {
	Type      string
	Amount    float32
	StartDate time.Time
	Status    string
}
