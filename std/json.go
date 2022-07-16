package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

type Order struct {
	ID          string        `json:"id"`
	DateOrdered time.Time     `json:"date_ordered"`
	CustomerID  string        `json:"customer_id"`
	Items       []interface{} `json:"items"`
}

type OrderController struct {
	mu          sync.Mutex
	countrycode string
}

func (oc *OrderController) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		rw.WriteHeader(400)
		rw.Write([]byte{})
		return
	}

	buf := make([]byte, 2048)

	n, err := r.Body.Read(buf)

	fmt.Println(n)

	if err != nil && err != io.EOF {
		log.Println(err)
		rw.WriteHeader(500)
		return
	}

	defer r.Body.Close()

	var o Order

	fmt.Println(buf[:n])

	errJson := json.Unmarshal(buf[:n], &o)

	if errJson != nil {
		log.Println(errJson)
		rw.WriteHeader(500)
		return
	}

	log.Println(o)
	res, err := json.Marshal(&o)

	if err != nil {
		log.Println(err)
		rw.WriteHeader(500)
		return
	}

	rw.Write(res)

}

func main() {

	http.Handle("/order", &OrderController{})
	http.ListenAndServe(":8080", nil)
}
