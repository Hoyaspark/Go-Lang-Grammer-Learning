package main

import (
	"fmt"
	"time"
)

func main() {
	duration := (2 * time.Hour) + (30 * time.Minute)
	fmt.Println(duration)
	time.AfterFunc(time.Second*3, func() {
		fmt.Println("time.AfterFunc has close()")
	})

	time.Sleep(time.Second * 10)
}
