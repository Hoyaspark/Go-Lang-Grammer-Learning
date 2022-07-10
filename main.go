package main

import (
	"fmt"
	"reflect"
)

func main() {
	//array()
	//slice()
	//arrayDifSlice()
	//string()
	sliceAppend()
	makeSlice()
}

func array() {
	var x = [...]int{1, 2, 3}
	var y = [...]int{1, 2, 3}
	fmt.Println(x == y)
	fmt.Println(reflect.TypeOf(x))
}

func slice() {
	var x = []int{10, 20, 30}
	var y []int

	// slice 는 == 연산이 불가능하다
	// fmt.Println(x == y)
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.TypeOf(y))

	l := len(y)

	fmt.Println(l)
	xx := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println(xx)
}

func sliceAppend() {
	var x []int

	x = append(x, 3)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 3)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 3)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 3)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 3)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 3)
	fmt.Println(x, len(x), cap(x))
}

func makeSlice() {
	// cap을 지정하는 이유는 slice에 값을 추가할때마다 할당되는 메모리 사이즈가 두배로 늘어나기 때문에 이를 제한하고자 cap을 일정하게 지정한다.
	x := make([]int, 0, 10) // type, len, cap
	y := make(map[string]string)

	// {} <- 이게 literal(리터럴)이라 부른다
	z := []int{}

	fmt.Println(z, len(z), cap(z))

	z = append(z, 3)
	fmt.Println(z, len(z), cap(z))

	y["hello"] = "gello"

	x = append(x, 3)
	fmt.Println(x)
	fmt.Println(y)

}

func arrayDifSlice() {
	var y []int

	if y == nil {
		fmt.Println("slice is nil")
	}
}
