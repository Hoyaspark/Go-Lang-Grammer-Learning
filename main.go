package main

import "fmt"

func main() {
	//array()
	//slice()
	//arrayDifSlice()
	//string()
	//sliceAppend()
	//makeSlice()
	//sliceMemoryShare()
	//conflictSlice()
	//arrayToSlice()
	//deepCopySlice()
	//stringSlice()

	//createMap()
	//findMap()
	//deleteMap()

	//basicStruct()
	//anonymousStruct()
	//getFor()
	//test()
	hello()
}

func hello() {
	var result []string

	result = append(result, "abcd")

	fmt.Println(result, len(result), cap(result))
}
