package main

import (
	"math/rand"
	"net"
	"sync"

	"github.com/andrewarrow/mini/lib"
)

var Last100Posts = []lib.MiniPost{}
var Last1000Posts = map[string]lib.MiniPost{}
var Mutex100 sync.Mutex
var Mutex1000 sync.Mutex

func ListenForPosts() {
	go func() {
		for mp := range lib.MiniPostChan {
			if mp.Body == "" {
				continue
			}
			Mutex100.Lock()
			Last100Posts = append([]lib.MiniPost{mp}, Last100Posts...)
			if len(Last100Posts) > 100 {
				Last100Posts = Last100Posts[0 : len(Last100Posts)-1]
			}
			Mutex100.Unlock()

			Mutex1000.Lock()
			if len(mp.ImageURLs) > 0 {
				Last1000Posts[mp.PostHashHex] = mp
			}
			if len(Last1000Posts) > 1000 {
				ranIndex := rand.Intn(1000)
				i := 0
				key := ""
				for k, _ := range Last1000Posts {
					if i == ranIndex {
						key = k
						break
					}
					i++
				}
				delete(Last1000Posts, key)
			}
			Mutex1000.Unlock()
		}
	}()

	lib.Connect("peer1", net.ParseIP("35.232.92.5"))
}
