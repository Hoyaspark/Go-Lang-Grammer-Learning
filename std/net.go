package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

// 기본 http.Get,http.Post 등을 피하자. 이유로 해당 기본 코드는 timeout을 설정할 수 없기 때문이다.
func requestServer() {
	client := http.Client{Timeout: 30 * time.Second}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.naver.com", nil)

	if err != nil {
		log.Println(err)
		return
	}
	res, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return
	}

	for {
		buf := make([]byte, 2048)
		n, err := res.Body.Read(buf)

		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(string(buf[:n]))
	}
}

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte(`Hello, My Name is
Pedro`))
}

func makeServer() {
	mux := http.NewServeMux()

	handler := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("hello, worlds!"))
	})

	mux.HandleFunc("/", handler)

	s := &http.Server{
		Addr:         ":80",
		Handler:      mux,
		ReadTimeout:  0,
		WriteTimeout: 0,
		IdleTimeout:  0,
	}

	err := s.ListenAndServe()

	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}

}

func makeMiddleWare() {
	mux := http.NewServeMux()

	s := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	tsp := TerribleSecurityProvider("GOPHER")

	mux.Handle("/auth", tsp(RequestTimer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello, worlds"))
	}))))

	err := s.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func RequestTimer(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(rw, r)
		end := time.Now()
		log.Printf("request time for %s: %v", r.URL.Path, end.Sub(start))
	})
}

func TerribleSecurityProvider(password string) func(handler http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			if r.Header.Get("X-Secret-Password") != password {
				rw.WriteHeader(http.StatusUnauthorized)
				rw.Write([]byte("You didn't give the secret passowrd\n"))
				return
			}
			h.ServeHTTP(rw, r)
		})
	}
}

func main() {
	//requestServer()
	//makeServer()
	makeMiddleWare()
}
