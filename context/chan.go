package main

import (
	"fmt"
	"time"
)

type key int

const (
	userKey key = iota
	boardKey
	famousKey
)

func main() {
	ch := make(chan struct{})

	go func() {
		time.Sleep(time.Second * 5)
		close(ch)
	}()

	for {
		select {
		case <-ch:
			fmt.Println("hello")
		}
	}

}
