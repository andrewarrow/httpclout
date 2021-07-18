package main

import (
	"net"
	"sync"

	"github.com/andrewarrow/mini/lib"
)

var Last100Posts = []lib.MiniPost{}
var Mutex sync.Mutex

func ListenForPosts() {
	go func() {
		for mp := range lib.MiniPostChan {
			if mp.Body == "" {
				continue
			}
			Mutex.Lock()
			Last100Posts = append([]lib.MiniPost{mp}, Last100Posts...)
			if len(Last100Posts) > 100 {
				Last100Posts = Last100Posts[0 : len(Last100Posts)-1]
			}
			Mutex.Unlock()
		}
	}()

	lib.Connect("peer1", net.ParseIP("35.232.92.5"))
}
