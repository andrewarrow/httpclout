package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/andrewarrow/cloutcli"
	"github.com/gin-gonic/gin"
	"github.com/justincampbell/timeago"
)

func WelcomeIndex(c *gin.Context) {
	username := c.Query("username")
	list := cloutcli.FollowingFeedPosts(username)
	fmt.Println(list)
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"posts": list,
	})
}

func RoutesSetup(router *gin.Engine) {

	//router.Static("/static", "static")
	router.GET("/", WelcomeIndex)

	AddTemplates(router)
}

func AddTemplates(r *gin.Engine) {
	fm := template.FuncMap{
		"mod": func(i, j int) bool { return i%j == 0 },
		"ago": func(i int64) string {
			d, _ := time.ParseDuration(fmt.Sprintf("%ds", time.Now().Unix()-i))
			return timeago.FromDuration(d)
		},
		"tokens": func(s string, i int) string { return strings.Split(s, ".")[i] },
		"add":    func(i, j int) int { return i + j },
	}
	r.SetFuncMap(fm)
	r.LoadHTMLGlob("templates/*.tmpl")
}
