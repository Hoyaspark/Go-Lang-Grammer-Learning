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

	pp := newPer()

	pp.updateAge(13)

	fmt.Println(pp)

	(*pp).failedUpdateAge(10)

	fmt.Println(pp)

	var m map[string]string

	// map은 nil값을 가질 수 있다
	if m == nil {
		fmt.Println("map is nil")
	}

	var ss []string

	// slice도 nil값을 가질 수 있다
	if ss == nil {
		fmt.Println("slice is nil")
	}

	//var st struct{}
	//
	//// struct는 nil값을 가질 수 없다.
	//if st == nil {
	//
	//}

	var st *struct{}

	// 하지만 *struct는 nil값을 가질 수 있다.
	if st == nil {
		fmt.Println("struct pointer is nil")
	}

	temp := pe{
		FirstName: "a",
		LastName:  "b",
		age:       1,
	}

	fu := temp.failedUpdateAge

	fu(13)

	fus := pe.failedUpdateAge

	fus(temp, 13)

	newPer().stream().stream().stream().stream().stream()

	(&temp).failedUpdateAge(13)

	fmt.Println(temp)

}

func newPer() *pe {
	return &pe{
		FirstName: "sang ho",
		LastName:  "park",
		age:       18,
	}
}

func (p *pe) updateAge(age int) {
	p.age = age // p->age = age
}

func (p pe) failedUpdateAge(age int) {
	fmt.Println(p)
	p.age = age
}

func (p pe) controlNil() {
	// 구조체 포인터는 nil 처리가 가능하지만 그냥 구조체는 nil이 먹히지 않는다!
	//if p == nil {
	//
	//}
}

func (p *pe) stream() *pe {
	return p
}

func conv(str string) score {
	return 3
}
