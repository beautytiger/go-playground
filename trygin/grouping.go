package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}

	router.Run(":8080")
}

func loginEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}

func submitEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}

func readEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}
