package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")

	if err != nil {
		log.Println(err)
	}

	fmt.Println(listener)

	http.Serve(listener, http.HandlerFunc(hello))
	// goroutine 하나 생성
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello,worlds"))
}
