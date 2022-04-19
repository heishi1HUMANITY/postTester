package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func handler(c *gin.Context) {
	fmt.Printf("access from %s\n", c.Request.RemoteAddr)
	fmt.Println("Header: ")
	for key, value := range c.Request.Header {
		fmt.Printf("%s: %s\n", key, value)
	}
	buf := make([]byte, 2048)
	n, _ := c.Request.Body.Read(buf)
	body := string(buf[:n])
	fmt.Printf("body: %s", body)
	c.Status(200)
}

func main() {
	server := gin.Default()
	server.GET("/", handler)
	server.POST("/", handler)
	server.Run(":8080")
}
