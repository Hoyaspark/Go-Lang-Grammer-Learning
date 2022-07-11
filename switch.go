package main

import (
	"fmt"
	"math/rand"
)

func main() {
	words := []string{"hi", "salutations", "hello"}

	for _, word := range words {
		switch exactWord := len(word + "2"); {
		case exactWord > 5:
			fmt.Println("hello")
		case exactWord < 5:
			fmt.Println("this")
		}
	}

	for i := 1; i < 10; i++ {
		switch num := rand.Intn(i); num {
		case 1, 2, 3, 4, 5:
			fmt.Println(1)
		default:
			fmt.Println(2)

		}
	}

}
