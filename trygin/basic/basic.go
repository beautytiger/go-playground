package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)


func main() {
	r := gin.Default()

	r.GET("/", pingPong)
	r.GET("/c", clientCancel)

	r.Run(":8080")
}

func pingPong(c *gin.Context) {
	time.Sleep(time.Second*10)
	c.String(http.StatusOK, "pong")
}

func clientCancel(c *gin.Context) {
	t := time.NewTicker(time.Second)
	defer t.Stop()
	for i := 0; i < 15; i++ {
		select {
		case <- t.C:
			log.Println("continue", i)
		case <- c.Request.Context().Done():
			log.Println("request cancel")
			return
		}
	}
	c.JSON(200, gin.H{"msg": "pong"})
}