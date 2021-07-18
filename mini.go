package main

import (
	"net"
	"sync"

	"github.com/andrewarrow/mini/lib"
)

var Last100Posts = []lib.MiniPost{}
var Last1000Posts = []lib.MiniPost{}
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
			if mp.PosterPub58 != "BC1YLfg6rAXxDdcJ95WRe9kKEbqfEgih4ewad1oXHbt4CKk2Mx22e5n" {
				Last1000Posts = append([]lib.MiniPost{mp}, Last1000Posts...)
			}
			if len(Last1000Posts) > 1000 {
				Last1000Posts = Last1000Posts[0 : len(Last1000Posts)-1]
			}
			Mutex1000.Unlock()
		}
	}()

	lib.Connect("peer1", net.ParseIP("35.232.92.5"))
}
