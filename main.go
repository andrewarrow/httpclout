package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	RoutesSetup(router)
	router.Run()
}
