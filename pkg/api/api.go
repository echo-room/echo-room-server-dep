package api

import (
	"github.com/gin-gonic/gin"
)

// Start starts the  API server
func Start() {
	router := gin.Default()
	router.GET("/", hello)
	router.Run(":3000")
}

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello back to you!",
	})
}
