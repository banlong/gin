package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	admin := r.Group("/admin")
	admin.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin-overview.html", nil)
	})

	//Using Relative path, map path "/public" to public folder "./public"
	r.Static("/public", "./public")

	//Using file Static file system, same as above
	//r.StaticFS("/public", http.Dir("./public"))

	r.Run(":3000")
}

