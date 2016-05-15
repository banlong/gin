package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	//Default has middleware integrated, can use gin.New() if you wnat to start from scratch
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello from %v", "Gin")
	})
	r.Run(":3000")
}