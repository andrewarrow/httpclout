package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BiggestFanOfShow(c *gin.Context) {
	username := c.Param("username")
	likes := []string{"wfwe"}
	reclouts := []string{"wfwefew", "Wfwe"}
	diamonds := []string{"wfwe"}
	c.HTML(http.StatusOK, "biggest_fan_of.tmpl",
		gin.H{"username": username,
			"likes": likes, "reclouts": reclouts, "diamonds": diamonds})
}
