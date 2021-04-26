package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Name": "Chen",
			"Sex":  "Male",
		})
	})
	r.Run(":8000")
}
