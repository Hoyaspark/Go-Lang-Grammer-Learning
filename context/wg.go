package main

import (
	"sync"
	"time"
)

func main() {
	//var wgValue sync.WaitGroup
	//
	//wgValue.Add(1)
	//go copyValue(wgValue)
	//wgValue.Wait()

	var wgPointer sync.WaitGroup

	wgPointer.Add(1)
	go copyPointer(&wgPointer)
	wgPointer.Wait()

}

func copyValue(wg sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(time.Second * 3)
}

func copyPointer(wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(time.Second * 3)
}
