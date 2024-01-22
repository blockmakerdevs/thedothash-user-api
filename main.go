package main

import (
	"github.com/gin-gonic/gin"
	redsRef "main/Config/redis"
)

func main() {
	redsRef.ConnectToRedis()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		err := redsRef.Client.Set("mykey", c.Query("value"), 0).Err()
		if err != nil {
			panic(err)
		}

		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	r.GET("/pong", func(c *gin.Context) {

		val, error := redsRef.Client.Get("mykey").Result()
		if error != nil {
			panic(error)
		}

		c.JSON(200, gin.H{
			"message": val,
		})
	})
	r.Run()
}
