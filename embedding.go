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

type Inner struct {
	A int
}

func (I Inner) IntPrinter(val int) string {
	return fmt.Sprintf("Inner : %d\n", val)
}

func (I Inner) Double() string {
	return I.IntPrinter(I.A * 2)
}

type Outer struct {
	Inner
	S string
}

func (o Outer) IntPrinter(val int) string {
	return fmt.Sprintf("Outer : %d\n", val)
}

func main() {
	o := Outer{
		Inner{A: 10},
		"Hello",
	}

	fmt.Println(o.Double())

}
