package main

import "fmt"

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

}
