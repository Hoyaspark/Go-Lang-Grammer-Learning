package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(1)

	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
		fmt.Println("send", i*2)
	}

	wg.Wait()

}

func square(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second)
	terminate := time.After(time.Second * 10)

	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-terminate:
			fmt.Println("terminate")
			wg.Done()
			return
		case v := <-ch:
			fmt.Println("received data", v)
		}
	}

}
