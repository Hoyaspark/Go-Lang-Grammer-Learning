package main

import (
	"fmt"
)

type per struct {
	firstName string
	lastName  *string
	age       int
}

func main() {

	p := per{
		firstName: "park",
		lastName:  getString("sangho"),
		age:       18,
	}

	fmt.Println(p)

	var f *int
	failedUpdate(f)

	fmt.Println(f)

	x := 10
	failedUpdate(&x)
	fmt.Println(x == 10)

	update(&x)
	fmt.Println(x == 20)

	failedUpdateStruct(p)
	fmt.Println(p)

	updateStruct(&p)
	fmt.Println(p)

}

func failedUpdateStruct(arg per) {
	arg.firstName = "pedro"
}

func updateStruct(arg *per) {
	arg.firstName = "hg"
}

func update(g *int) {
	*g = 20
}

func failedUpdate(g *int) {
	a := 3

	g = &a
}

func getString(str string) *string {
	return &str
}
