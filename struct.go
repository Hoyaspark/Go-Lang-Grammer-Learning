package main

import "fmt"

type person struct {
	name string
	age  int
	pet  string
}

func (p *person) String() string {
	v := p.name + " has " + p.pet + " & age is " + string(p.age)

	return v
}

func createNewPerson() *person {
	return &person{
		name: "pedro",
		age:  18,
		pet:  "pd",
	}
}

func basicStruct() {
	p := createNewPerson()

	fmt.Println(p)
}

func anonymousStruct() {
	g := struct {
		name string
		age  int
		pet  string
	}{
		name: "pedro",
		age:  18,
		pet:  "pd",
	}

	fmt.Println(g == person{
		name: "pedro",
		age:  18,
		pet:  "pd",
	})

	src := person{
		name: "pedro",
		age:  18,
		pet:  "pd",
	}

	dest := person{
		name: "pedro",
		age:  18,
		pet:  "pd",
	}

	fmt.Println(src == dest)

	//srcS := []string{"a"}
	//dstS := []string{"a"}

	// slice can only be compared to nil!!
	//fmt.Println(srcS == dstS)
}
