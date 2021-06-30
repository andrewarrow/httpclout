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
	path := "/home/aa/httpclout/user_sqlites/"
	likes := cloutcli.QuerySqliteTopLikers(path+username+".db", "50")
	reclouts := cloutcli.QuerySqliteTopReclouters(path+username+".db", "50")
	diamonds := cloutcli.QuerySqliteTopDiamondGivers(path+username+".db", "50")
	//usernameMutex.Unlock()
	c.HTML(http.StatusOK, "biggest_fan_of.tmpl",
		gin.H{"username": username,
			"likes": likes, "reclouts": reclouts, "diamonds": diamonds})
}
