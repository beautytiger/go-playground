package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	s := gin.Default()
	h := new(handler)
	s.POST("/", h.printName)
	s.Run(":8000")
}

// 使用这种handler进行直接json绑定，会导致所有请求复用同一个对象，后果很严重且不可知
type handler struct {
	Name string
}

func (h *handler)printName(c *gin.Context) {
	log.Printf("%p", h)
	c.BindJSON(h)
	init := h.Name
	for i := 0; i < 100; i++ {
		fmt.Printf("init: %s, current: %s\n", init, h.Name)
		time.Sleep(1*time.Second)
	}
	c.JSON(200, gin.H{})
}
