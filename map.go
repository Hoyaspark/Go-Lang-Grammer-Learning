package main

import "fmt"

func createMap() {
	m := make(map[string]string, 5)

	fmt.Println(len(m))
}

func findMap() {
	m := make(map[string]string)

	m["name"] = "pedro"

	// if present exists is true
	value, exists := m["name"]

	fmt.Println(value, exists)

	v, e := m["hello"]

	// if not present exists is false, value is ""
	fmt.Println(v == "", e)

}

func deleteMap() {
	m := map[string]string{
		"name": "pedro",
		"age":  "18",
	}

	delete(m, "name")
	_, exists := m["name"]
	fmt.Println(exists, "is must false")

	delete(m, "name")
	// nothing happen

	if v, e := m["age"]; e {
		fmt.Println(v)
	}
}
