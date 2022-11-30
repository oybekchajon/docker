package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message":"pong",
		})
	})
	r.GET("/pong", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message":"ping",
		})
	})
	r.Run(":8080")
}