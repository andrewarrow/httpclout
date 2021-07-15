package main

import (
	"net"

	"github.com/andrewarrow/mini/lib"
)

var Last100Posts = []lib.MiniPost{}

func ListenForPosts() {
	go func() {
		for mp := range lib.MiniPostChan {
			if mp.Body == "" {
				continue
			}
			Last100Posts = append([]lib.MiniPost{mp}, Last100Posts...)
			if len(Last100Posts) > 100 {
				Last100Posts = Last100Posts[0 : len(Last100Posts)-1]
			}
		}
	}()

	lib.Connect("peer1", net.ParseIP("35.232.92.5"))
}
