package main

import (
	"fmt"
	"reflect"
)

/**
기억하자 GoLang은 Call By Value이다.
*/
func main() {

	dump := []struct {
		name string
		age  int
	}{
		{name: "pedro", age: 18},
		{name: "hoyas", age: 25},
	}

	for _, value := range dump {
		fmt.Println(value.name, "is", value.age)
	}

	m := map[string]bool{
		"name":  true,
		"hello": false,
	}

	for k, v := range m {
		fmt.Println(k, ":", v)
		v = false
	}

	fmt.Println(m)

	samples := []string{"hello", "apple_😄!"}

	// 문자열을 range로 돌때는 rune (int32)형식을 사용한다
	// offset은 int8 UTF-8 형식에 따라간다.
	for _, sample := range samples {
		fmt.Println([]byte(sample))
		fmt.Println([]rune(sample))
		for i, r := range sample {
			fmt.Println(i, r, string(r), reflect.TypeOf(r))
		}
		fmt.Println()
	}

	// 그냥 완전한 for문을 이용하여 byte단위로 접근해보자.
	for _, sample := range samples {
		for i := 0; i < len(sample); i++ {
			fmt.Println(i, sample[i], string(sample[i]), reflect.TypeOf(sample[i]))
		}

		fmt.Println()
	}

}
