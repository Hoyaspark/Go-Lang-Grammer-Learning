package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("no possible arg")
	}

	f, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatalln("error open file")
	}

	defer f.Close()

	b := make([]byte, 128)

	for {
		_, err := f.Read(b)

		os.Stdout.Write(b)

		if err != nil {
			if err == io.EOF {
				log.Println("End Of File...")
				return
			}
		}
	}

	defer func() {
		callMessage()
	}()
}

func callMessage() {
	fmt.Println("Hello")
}
