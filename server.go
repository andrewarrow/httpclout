package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/andrewarrow/cloutcli"
	"github.com/andrewarrow/cloutcli/lib"
	"github.com/gin-gonic/gin"
	"github.com/justincampbell/timeago"
)

func WelcomeIndex(c *gin.Context) {
	username := c.Query("username")
	exclude := c.Query("exclude")
	excludeList := strings.Split(exclude, ",")
	list := []lib.Post{}
	for _, item := range cloutcli.FollowingFeedPosts(username) {
		skip := false
		for _, x := range excludeList {
			if x != "" && strings.Contains(item.Body, x) {
				skip = true
				break
			}
		}
		if skip {
			continue
		}
		list = append(list, item)
	}

	for _, item := range list {
		item.TimestampNanos = item.TimestampNanos / 1000000000
	}
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
