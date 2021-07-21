package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := 9876

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	_ = r.Run(fmt.Sprintf("localhost:%d", port))


	os.Exit(1)
}
