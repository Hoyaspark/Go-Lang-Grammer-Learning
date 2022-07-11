package main

import (
	"fmt"
)

type score int
type converter func(string) score
type teamScore map[string]score

type pe struct {
	FirstName string
	LastName  string
	age       int
}

func main() {

	var s score = 3

	var c converter = conv

	t := teamScore{}

	fmt.Println(s, c, t)

	p := pe{
		FirstName: "park",
		LastName:  "sang ho",
		age:       25,
	}

	fmt.Println(p)

	p.updateAge(18) // 자동적으로 (&p).updateAge(18)로 변환된다

	fmt.Println(p)

	p.failedUpdateAge(22)

	fmt.Println(p)
}

func (p *pe) updateAge(age int) {
	p.age = age // p->age = age
}

func (p pe) failedUpdateAge(age int) {
	p.age = age
}

func conv(str string) score {
	return 3
}
