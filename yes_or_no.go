package main

import (
	"math/rand"
	"net/http"

	"github.com/andrewarrow/mini/lib"
	"github.com/gin-gonic/gin"
)

func YesOrNoIndex(c *gin.Context) {
	pub58, _ := c.Cookie("identity_pub58")
	Mutex1000.Lock()
	if len(Last1000Posts) == 0 {
		Mutex1000.Unlock()
		c.String(http.StatusOK, "try again please")
		return
	}
	ranIndex := rand.Intn(len(Last1000Posts))
	i := 0
	ranPost := lib.MiniPost{}
	for _, v := range Last1000Posts {
		if i == ranIndex {
			ranPost = v
			break
		}
		i++
	}
	Mutex1000.Unlock()
	//ranPost := lib.MiniPost{}
	//ranPost.PostHashHex = "9ef3afa082054898b21f056211e2c6f3f145bdea62d9805eb9e59e1538d595dc"
	//ranPost.PosterPub58 = "BC1YLiCo6prb6M3xELpRbHUAtQvNAegcr2GHg1Z9LYDL52cZrbctHmr"

	successPostHashHex := c.Query("postHashHex")
	successTheirPub58 := c.Query("theirPub58")
	c.HTML(http.StatusOK, "yes_or_no.tmpl", gin.H{"successPostHashHex": successPostHashHex,
		"successTheirPub58": successTheirPub58,
		"pub58":             pub58, "Post": ranPost})
}
