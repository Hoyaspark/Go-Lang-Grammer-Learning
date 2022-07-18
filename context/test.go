package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)

	defer cancel()

	var wg sync.WaitGroup

	wg.Add(2)

	go a(ctx, &wg)
	go b(ctx, &wg)

	wg.Wait()

}

func a(ctx context.Context, wg *sync.WaitGroup) {
	tick := time.Tick(time.Second)

	context, cancel := context.WithCancel(ctx)

	defer cancel()

	go ab(context)
	go ac(context)

	for {
		select {
		case <-tick:
			fmt.Println("tick from func A")
		case <-context.Done():
			fmt.Println("done func A")
			wg.Done()
			return
		}
	}
}

func b(ctx context.Context, wg *sync.WaitGroup) {
	tick := time.Tick(time.Second)

	context, cancel := context.WithCancel(ctx)

	defer cancel()

	for {
		select {
		case <-tick:
			fmt.Println("tick from func B")
		case <-context.Done():
			fmt.Println("done func B")
			wg.Done()
			return
		}
	}
}

func ab(ctx context.Context) {
	tick := time.Tick(time.Second)

	context, cancel := context.WithCancel(ctx)

	defer cancel()

	go abc(context)
	go abd(context)

	for {
		select {
		case <-tick:
			fmt.Println("tick from func AB")
		case <-context.Done():
			fmt.Println("done func AB")
			return
		}
	}
}

func ac(ctx context.Context) {
	tick := time.Tick(time.Second)

	for {
		select {
		case <-tick:
			fmt.Println("tick from func AC")
		case <-ctx.Done():
			fmt.Println("done func AC")
			return
		}
	}
}

func abc(ctx context.Context) {
	tick := time.Tick(time.Second)

	context, cancel := context.WithCancel(ctx)

	defer cancel()

	for {
		select {
		case <-tick:
			fmt.Println("tick from func ABC")
		case <-context.Done():
			fmt.Println("done func ABC")
			return
		}
	}
}

func abd(ctx context.Context) {
	tick := time.Tick(time.Second)
	context, cancel := context.WithCancel(ctx)

	defer cancel()

	for {
		select {
		case <-tick:
			fmt.Println("tick from func ABD")
		case <-context.Done():
			fmt.Println("done func ABD")
			return
		}
	}
}
