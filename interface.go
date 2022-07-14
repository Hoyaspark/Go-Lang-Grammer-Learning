package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

type LogicProvider struct{}

func (lp LogicProvider) Process(data string) string {
	return data
}

type Logic interface {
	Process(data string) string
}

type LogicTest struct {
	Type reflect.Type
}

func (l *LogicTest) test() {
	l.Type.Kind()
}

type ttt int

func (t *ttt) test() *int {
	cc := int(*t)
	d := &cc

	*d = 4
	return d
}

type handlerFunc func(http.ResponseWriter, *http.Request)

func (f handlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

type Client struct {
	L Logic
}

func (c Client) Program() {
	fmt.Println(c.L.Process("pedro"))
}

type MyInt int

func PrintVal(v interface{}) {

	switch a := 3; a {
	case 3:
		fmt.Println("3")

	}
}

func main() {

	c := Client{
		L: LogicProvider{},
	}

	c.Program()

	data := map[string]interface{}{}

	contents, err := ioutil.ReadFile("test.json")

	if err != nil {
		log.Fatalln(err)
	}

	err2 := json.Unmarshal(contents, &data)

	for k, v := range data {
		fmt.Println(k)

		switch v := v.(type) {
		case string:
			fmt.Println("string :", v)
		}
	}

	if err2 != nil {
		log.Fatalln(err2)
	}

	fmt.Println(data)

	//switch j := a.(type); j {
	//case nil:
	//	fmt.Println("hello")
	//default:
	//	fmt.Println("hello")
	//
	//}
	//var h handlerFunc
	//h = func(writer http.ResponseWriter, request *http.Request) {
	//	switch j := writer.(type);j {
	//	case
	//
	//	}
	//}
	//h.ServeHTTP(nil, nil)
}
