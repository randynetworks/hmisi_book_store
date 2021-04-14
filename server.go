package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"hmisiBookStore/conf"
)
func main() {


	route := gin.Default()
	route.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Halo selamat datang di RESTAPI HMISI Book Store",
		})
	})
	_ = route.Run() // listen and serve on 0.0.0.0:8080
}
