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

func RoutesSetup(router *gin.Engine) {

	//router.Static("/static", "static")
	router.GET("/", WelcomeIndex)
	router.GET("/exclude", ExcludeIndex)
	router.GET("/biggest-fans-of/:username", BiggestFanOfShow)

	AddTemplates(router)
}

func WelcomeIndex(c *gin.Context) {
	//httpclout_cookie1
	pub58, _ := c.Cookie("httpclout_pub58")
	if pub58 == "" {
		c.HTML(http.StatusOK, "welcome.tmpl", gin.H{})
	} else {
		c.HTML(http.StatusOK, "feed.tmpl",
			gin.H{"baseURL": "http://192.168.1.50:17001", "pub58": pub58})
	}
	return
}
func ExcludeIndex(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		c.HTML(http.StatusOK, "form.tmpl", gin.H{})
		return
	}
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
