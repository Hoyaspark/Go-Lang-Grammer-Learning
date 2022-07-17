package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Order struct {
	ID          string        `json:"id"`
	DateOrdered RFC822ZTime   `json:"date_ordered"`
	CustomerID  string        `json:"customer_id"`
	Items       []interface{} `json:"items"`
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type TestController struct {
}

func (tc *TestController) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var buf Order

	err := json.NewDecoder(r.Body).Decode(&buf)

	fmt.Println(buf)

	if err != nil {
		log.Println(err)
		rw.WriteHeader(500)
		return
	}

	defer r.Body.Close()

	json.NewEncoder(rw).Encode(&buf)

	fmt.Println(buf)

	//d := json.NewDecoder(r.Body)
	//
	//var bufArray []interface{}
	//for d.More() {
	//	var buf interface{}
	//	d.Decode(&buf)
	//
	//	bufArray = append(bufArray, buf)
	//}
	//
	//fmt.Println(bufArray)

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

	buf := make([]byte, 30)

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

type Message struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

func jsonStream() ([]Message, error) {
	const testJson string = `
	[
		{"name": "Ed", "text": "Knock knock."},
		{"name": "Sam", "text": "Who's there?"},
		{"name": "Ed", "text": "Go fmt."},
		{"name": "Sam", "text": "Go fmt who?"},
		{"name": "Ed", "text": "Go fmt yourself!"}
	]
`

	d := json.NewDecoder(strings.NewReader(testJson))

	var result []Message

	t, err := d.Token()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	fmt.Println(t)

	for d.More() {
		var buf Message
		err = d.Decode(&buf)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		result = append(result, buf)
	}

	t, err = d.Token()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	fmt.Println(t)

	return result, nil
}

//Custom Json Parsing
type RFC822ZTime struct {
	time.Time
}

func (rt RFC822ZTime) MarshalJSON() ([]byte, error) {
	out := rt.Time.Format(time.RFC822)

	return []byte("\"" + out + "\""), nil
}

func (rt *RFC822ZTime) UnmarshalJSON(b []byte) error {
	t, err := time.Parse("\""+time.RFC822+"\"", string(b))

	if err != nil {
		log.Println(err)
		return err
	}

	rt = &RFC822ZTime{t}

	return nil
}

func customJSONParse() ([]Order, error) {
	const jsonText = `
[
	{
		"id":"12345",
		"date_ordered":"2020-05-01T13:01:02Z",
		"customer_id":"3",
		"items":[{"id":"pedro","name":"Here"}]
	},
	{
		"id":"12345",
		"date_ordered":"2020-05-01T13:01:02Z",
		"customer_id":"3",
		"items":[{"id":"pedro","name":"Here"}]
	},
	{
		"id":"12345",
		"date_ordered":"2020-05-01T13:01:02Z",
		"customer_id":"3",
		"items":[{"id":"pedro","name":"Here"}]
	},
	{
		"id":"12345",
		"date_ordered":"2020-05-01T13:01:02Z",
		"customer_id":"3",
		"items":[{"id":"pedro","name":"Here"}]
	}
]
`
	d := json.NewDecoder(strings.NewReader(jsonText))

	t, err := d.Token()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	fmt.Println("token :", t)

	var result []Order

	for d.More() {
		var buf Order
		err = d.Decode(&buf)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		result = append(result, buf)
	}

	t, err = d.Token()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	fmt.Println("token :", t)

	return result, nil

}

func main() {

	//toFile := &Person{"pedro", 25}
	//
	//tmpFile, err := ioutil.TempFile(os.TempDir(), "sample-")
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer os.Remove(tmpFile.Name())
	//
	//err = json.NewEncoder(tmpFile).Encode(toFile)
	//
	//dec := json.NewDecoder(strings.NewReader("hello"))
	//
	//dec.More()
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = tmpFile.Close()
	//
	//if err != nil {
	//	panic(err)
	//}

	//r, err := jsonStream()
	//
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//
	//fmt.Println(r)

	//tmp := "2020-05-01T13:01:02Z"
	//t, err := time.Parse(time.RFC3339, tmp)
	//
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//
	//s := t.Format(time.RFC822)
	//fmt.Println(s)
	http.Handle("/order", &TestController{})
	http.ListenAndServe(":8080", nil)
}
