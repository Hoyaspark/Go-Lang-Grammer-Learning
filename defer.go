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

	defer func() {
		callMessage()
	}()

	b := make([]byte, 128)

	for {
		count, err := f.Read(b)

		os.Stdout.Write(b[:count])

		if err != nil {
			if err == io.EOF {
				log.Println("End Of File...")
				return
			}
		}
	}

}

func callMessage() {
	fmt.Println("Hello")
}
