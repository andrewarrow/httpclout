package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func YesOrNoIndex(c *gin.Context) {
	pub58, _ := c.Cookie("httpclout_pub58")
	ranPost := Last100Posts[rand.Intn(100)]
	//ranPost := lib.MiniPost{}
	if pub58 == "" {
		c.HTML(http.StatusOK, "yes_or_no.tmpl", gin.H{"Post": ranPost})
	}
}
