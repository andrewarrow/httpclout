package main

import (
	"net/http"
	"sync"

	"github.com/andrewarrow/cloutcli"
	"github.com/gin-gonic/gin"
)

var usernameMutex sync.Mutex

func BiggestFanOfShow(c *gin.Context) {
	username := c.Param("username")
	//usernameMutex.Lock()
	likes := cloutcli.QuerySqliteTopLikers(username, "50")
	reclouts := cloutcli.QuerySqliteTopReclouters(username, "50")
	diamonds := cloutcli.QuerySqliteTopDiamondGivers(username, "50")
	//usernameMutex.Unlock()
	c.HTML(http.StatusOK, "biggest_fan_of.tmpl",
		gin.H{"username": username,
			"likes": likes, "reclouts": reclouts, "diamonds": diamonds})
}
