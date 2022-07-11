package main

import (
	"fmt"
	"reflect"
	"strings"
)

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

// slice는 메모리를 공유한다.
func sliceMemoryShare() {
	x := []int{1, 2, 3}

	y := x

	fmt.Println(x)
	fmt.Println(y)

	y[0] = 4
	fmt.Println(x)
	fmt.Println(y)
}

func conflictSlice() {
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)

	y := x[:2]
	z := x[2:]

	fmt.Println(x, y, z)

	// y의 cap은 부모 slice의 cap 에서 offset을 뺀 만큼이다
	fmt.Println(cap(x), cap(y), cap(z))

	y = append(y, 30, 40, 50)
	x = append(x, 60)
	z = append(z, 70)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

func arrayToSlice() {
	x := [4]int{1, 2, 3, 4}
	y := x[:2]
	z := x[2:]

	x[0] = 10

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	/**
	expected :
	x [10 2 3 4]
	y [10 2]
	z [3 4]

	*/
}

func deepCopySlice() {
	x := []int{1, 2, 3, 4}
	y := make([]int, 2)

	// copy함수는 len에 영향을 받는다
	num := copy(y, x)

	fmt.Println(num, y)
}

func makeSlice() {
	// cap을 지정하는 이유는 slice에 값을 추가할때마다 할당되는 메모리 사이즈가 두배로 늘어나기 때문에 이를 제한하고자 cap을 일정하게 지정한다.
	x := make([]int, 5, 10) // type, len, cap
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

func stringSlice() {
	// len함수는 길이가 아닌 byte의 길이이다.
	var s string = "Hello there"

	all := strings.ReplaceAll(s, "Hello", "Gello")

	fmt.Println(all)

	var b byte = s[2]

	fmt.Println(reflect.TypeOf(b))

}

func test() {
	x := []int{1, 2, 3}
	y := make([]int, 0, 5)

	z := append(y, x...)

	fmt.Println(z)
}
