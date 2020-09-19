package controllers

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", hello)
	r.GET("/:shortened", FindUrl)

	return r
}

func hello(c *gin.Context) {
	c.String(200, "Hello there")
}
