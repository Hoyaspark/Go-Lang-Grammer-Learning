package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)

	defer cancel()

	data, err := longRunningThingManager(ctx, "hello")

	fmt.Println(data, err)

}

func longRunningThingManager(ctx context.Context, data string) (string, error) {
	type wrapper struct {
		result string
		err    error
	}

	ch := make(chan *wrapper, 1)

	go func() {
		result, err := longRunningThing(ctx, data)
		ch <- &wrapper{result, err}
	}()

	select {
	case data := <-ch:
		fmt.Println("get data from chan")
		return data.result, data.err
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func longRunningThing(ctx context.Context, data string) (string, error) {

	tick := time.Tick(time.Second)
	after := time.After(time.Second * 10)

	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-ctx.Done():
			fmt.Println("error occurs")
			return "", ctx.Err()
		case <-after:
			fmt.Println("success done")
			return "succes", nil
		}
	}
}
