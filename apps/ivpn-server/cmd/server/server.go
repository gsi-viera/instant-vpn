package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "online"})
	})

	r.GET("/config", func(c *gin.Context) {
		key := c.Query("key")
		c.JSON(200, gin.H{
			"receivedKey": key,
		})
	})
	r.Run()
}
