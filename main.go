package main

import (
	"github.com/MgHtinLynn/configs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db := configs.Connection()

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
