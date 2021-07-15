package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	go ListenForPosts()
	router := gin.New()
	router.Use(gin.Logger())
	RoutesSetup(router)
	port := "3000"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	router.Run(":" + port)
}
