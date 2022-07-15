package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})

	go func() {
		time.Sleep(time.Second * 2)
		close(ch)
	}()

	select {
	case <-ch:
		fmt.Println("close(ch)")
		return
	}
}
