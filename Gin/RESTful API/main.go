package main

import "github.com/gin-gonic/gin"

type User struct {
	ID   uint64
	Name string
}

func main() {
	users := []User{{ID: 123, Name: "张三"}, {ID: 456, Name: "李四"}}
	r := gin.Default()
	r.DELETE("/users/123", func(c *gin.Context) {
		//删除ID为123的用户
		c.JSON(200, users)
	})
	r.Run(":8080")
}
