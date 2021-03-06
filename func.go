package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

type opFunc func(int, int) int

type a struct {
}

type bb []int

func main() {
	a := addTo(1, 2, 3, 4, 5, 6, 7, 8, 9)

	fmt.Println(a)

	b := []int{1, 2, 3, 4}

	c := addTo(-1, b...)

	fmt.Println(c)

	d := addTo(1)

	fmt.Println(d)

	fmt.Println()
	// correct
	div, rem, err := divAndRemainder(3, 2)

	checkErr(err)

	fmt.Println("div :", div, "rem :", rem)

	//// has Error
	//_, _, err2 := divAndRemainder(3, 0)
	//
	//checkErr(err2)

	r, e := availableTypeName(1, 2)

	fmt.Println(r, e)

	fmt.Println(mapToFunc(1, 2, '+'))
	fmt.Println(mapToFunc(1, 2, '-'))
	fmt.Println(mapToFunc(1, 2, '*'))
	fmt.Println(mapToFunc(1, 2, '/'))

	g := struct {
		name string
		age  int
	}{
		name: "pedro",
		age:  18,
	}

	fmt.Println(g)

}

// 가변 파라미터는 slice 로 변환된다
func addTo(base int, list ...int) []int {

	fmt.Println(reflect.TypeOf(list)) // []int

	a := make([]int, 0, len(list))
	for _, value := range list {
		a = append(a, base+value)
	}
	return a
}

func divAndRemainder(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}

	return numerator / denominator, numerator % denominator, nil
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// 빈 반환은 좋지 않다..
func availableTypeName(a, b int) (result int, err error) {
	return
}

func add(i, j int) int {
	return i + j
}

func sub(i, j int) int {
	return i - j
}

func mul(i, j int) int {
	return i * j
}

func div(i, j int) int {
	return i / j
}

func mapToFunc(i, j int, k rune) int {
	opMap := map[rune]func(int, int) int{
		'+': add,
		'-': sub,
		'*': mul,
		'/': div,
	}

	return opMap[k](i, j)
}

func anonymousFunc() {
	for i := 0; i < 5; i++ {
		// j에 i를 넘겨준다
		func(j int) {

		}(i)
	}
}
