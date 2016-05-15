package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//Default has middleware integrated, can use gin.New() if you wnat to start from scratch
	r := gin.Default()

	//Load all html files under the templates dir, including the sub-dir
	r.LoadHTMLGlob("templates/**/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	//Define Routes Group
	admin := r.Group("/admin")
	//handle /admin/
	admin.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin-overview.html", nil)
	})

	r.Run(":3000")
}
