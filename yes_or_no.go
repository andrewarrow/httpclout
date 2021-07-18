package main

import (
	"net/http"

	"github.com/andrewarrow/mini/lib"
	"github.com/gin-gonic/gin"
)

func YesOrNoIndex(c *gin.Context) {
	pub58, _ := c.Cookie("identity_pub58")
	//ranPost := Last100Posts[rand.Intn(100)]
	ranPost := lib.MiniPost{}
	c.HTML(http.StatusOK, "yes_or_no.tmpl", gin.H{"pub58": pub58, "Post": ranPost})
}
