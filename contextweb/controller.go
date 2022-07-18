package main

import (
	"net/http"
	"time"
)

type controller struct {
}

func (c *controller) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, ok := UserFromContext(ctx)

	if !ok {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := r.URL.Query().Get("data")
	// logic here retrieve data

	rw.Write([]byte(data + " " + user))

}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/user", Middleware(&controller{}))

	s := &http.Server{
		Addr:         "",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	s.ListenAndServe()

}
