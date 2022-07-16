package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := make(map[string]int)
	for {
		n, err := r.Read(buf)

		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}

		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}

func buildGzipReader(filename string) (*gzip.Reader, func(), error) {
	r, err := os.Open(filename)

	if err != nil {
		return nil, nil, err
	}

	gr, err := gzip.NewReader(r)

	if err != nil {
		return nil, nil, err
	}

	return gr, func() {
		gr.Close()
		r.Close()
	}, nil
}

func main() {
	s := "The quick brown fox jumped over the lazy dog"
	sr := strings.NewReader(s)

	out, err := countLetters(sr)

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(out)

	r, closer, err := buildGzipReader("my_data.txt.gz")
	if err != nil {
		log.Fatalln(err)
	}

	defer closer()

	counts, err := countLetters(r)

	fmt.Println(counts)

	//f, err := fileOpen("sdfsdf")
	//
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//fmt.Println(f)

	b, err := ioutil.ReadFile("hello")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(b)
}

func fileOpen(fileName string) (*os.File, error) {
	f, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	defer func() {
		fmt.Println("f is close()")
		f.Close()
	}()

	return f, nil
}
