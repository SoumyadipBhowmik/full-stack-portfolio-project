package main

import (
	"github.com/SoumyadipBhowmik/gba-landing/drivers"
	"github.com/gin-gonic/gin"
)

func init() {
	drivers.LoadEnvVariables()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
