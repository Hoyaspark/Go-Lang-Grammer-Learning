package main

import "fmt"

type Employee struct {
	name string
	id   string
}

type Manager struct {
	inta
	Employee
	reports []Employee
}

type inta int8

func main() {

	var s []Employee

	s = append(s, Employee{
		name: "sangho",
		id:   "dklfjdlfkjdflk",
	})

	m := Manager{
		inta: 3,
		Employee: Employee{
			name: "pedro",
			id:   "hello",
		},
		reports: s,
	}

	fmt.Println(m)
}
