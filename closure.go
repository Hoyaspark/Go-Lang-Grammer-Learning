package main

import (
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {

	people := []Person{
		{FirstName: "a", LastName: "b", Age: 1},
		{FirstName: "c", LastName: "d", Age: 5},
		{FirstName: "e", LastName: "f", Age: 3},
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)

}

func parameterSendFunc(arg any, less func(int, int) bool) {
}
