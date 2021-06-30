package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/http/httputil"
	u "net/url"
	"os"
	"strings"
	"time"

	"github.com/andrewarrow/cloutcli"
	"github.com/andrewarrow/cloutcli/lib"
	"github.com/andrewarrow/cloutcli/network"
	"github.com/gin-gonic/gin"
	"github.com/justincampbell/timeago"
)

func RoutesSetup(router *gin.Engine) {

	//router.Static("/static", "static")
	router.GET("/", WelcomeIndex)
	router.GET("/exclude", ExcludeIndex)
	router.GET("httpclout/biggest-fans-of/:username", BiggestFanOfShow)
	router.NoRoute(HandleApi)

	AddTemplates(router)
}

func HandleApi(c *gin.Context) {
	url, _ := u.Parse("http://localhost:17001")
	reverseProxy := httputil.NewSingleHostReverseProxy(url)
	reverseProxy.Director = func(req *http.Request) {
		req.Header.Add("X-Forwarded-Host", req.Host)
		req.URL.Scheme = url.Scheme
		req.URL.Host = url.Host
		req.Host = url.Host
	}

	reverseProxy.ServeHTTP(c.Writer, c.Request)
}
func WelcomeIndex(c *gin.Context) {
	pub58, _ := c.Cookie("httpclout_pub58")
	if pub58 == "" {
		c.HTML(http.StatusOK, "welcome.tmpl", gin.H{})
	} else {
		network.NodeURL = os.Getenv("CLOUT_API_INTERNAL_URL")
		items := cloutcli.FollowingFeedPub58(pub58)
		c.HTML(http.StatusOK, "feed.tmpl",
			gin.H{"baseURL": os.Getenv("CLOUT_API_EXTERNAL_URL"),
				"pub58": pub58, "items": items})
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
