package main

import "fmt"

type MailCategory int

const (
	Uncategorized MailCategory = iota
	Personal
	Spam
	Social
	Advertisements
)

func main() {

	fmt.Println(Uncategorized)
	fmt.Println(Personal)
	fmt.Println(Spam)
	fmt.Println(Social)
	fmt.Println(Advertisements)
}
