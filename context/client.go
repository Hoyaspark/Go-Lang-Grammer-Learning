package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var client http.Client

func callBoth(ctx context.Context, errVal, slowURL, fastURL string) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		err := callServer(ctx, "slow", slowURL)
		if err != nil {
			fmt.Println("slow context cancled()")
			cancel()
		}
	}()

	go func() {
		defer wg.Done()
		err := callServer(ctx, "fast", fastURL+"?error="+errVal)
		if err != nil {
			fmt.Println("fast context cancled()")
			cancel()
		}
	}()

	wg.Wait()
	fmt.Println("done with both")
}

func callServer(ctx context.Context, label, url string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		fmt.Println(label, "request err:", err)
		return err
	}

	res, err := client.Do(req)

	if err != nil {
		fmt.Println(label, "request err1:", err)
		return err
	}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(label, "request err2:", err)
		return err
	}

	defer res.Body.Close()

	result := string(data)

	if result != "" {
		fmt.Println(label, "result:", result)
	}

	if result == "error" {
		fmt.Println("cancelling from", label)
		return errors.New("error happened")
	}

	return nil

}
