package main

import (
	"fmt"
	"net"
	"time"

	"github.com/andrewarrow/mini/lib"
)

func ListenForPosts() {
	go func() {
		for mp := range lib.MiniPostChan {
			fmt.Println(mp.Body)
			fmt.Println("")
			fmt.Println(time.Unix(mp.Timestamp, 0))
			fmt.Println("")
			fmt.Println("https://bitclout.com/posts/" + mp.PostHashHex)
			fmt.Println("Poster Public Key", mp.PosterPub58)
			fmt.Println("")
			fmt.Println("")
		}
	}()

	lib.Connect("peer1", net.ParseIP("35.232.92.5"))
}
