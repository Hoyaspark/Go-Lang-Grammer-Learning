package main

import "fmt"

type golang struct {
	name string
	age  int
}

type lang golang

func main() {
	/**
	GoLang에서 func의 return은 복사일까?
	*/

	// 1. 포인터가 아닌 값을 리턴할때
	g := newGoLang()

	fmt.Println(g)

	fmt.Println()

	// 2. 포인터를 리턴할때
	gp := newAddressGoLang()

	fmt.Println(*gp)

}

func (g *golang) getName() string {
	return g.name
}

func newGoLang() golang {
	r := golang{
		name: "go",
		age:  2002,
	}

	fmt.Println(r)
	return r
}

func newAddressGoLang() *golang {
	r := golang{
		name: "go",
		age:  2002,
	}

	fmt.Println(r)

	return &r
}
