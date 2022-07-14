package main

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
)

type fmtError struct {
	message string
}

func (fe fmtError) Error() string {
	return fmt.Sprintf("error : %s\n", fe.message)
}

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

// StatusErr type is pedro
type StatusErr struct {
	status  Status
	message string
}

func (se StatusErr) Error() string {
	return se.message
}

func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
	if file == "abcd" {
		return nil, StatusErr{
			status:  InvalidLogin,
			message: "Invalid Login",
		}
	}

	return []byte("hello"), nil
}

func div60(i int) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()

	fmt.Println(300 / i)
}

func main() {
	fmt.Println(errors.Unwrap(fmtError{
		message: "hello",
	}))

	data := []byte("This is not a zip file")
	notAZipFile := bytes.NewReader(data)

	_, err := zip.NewReader(notAZipFile, int64(len(data)))

	if err == zip.ErrFormat {
		fmt.Println("Told you so")
	}

	for _, value := range []int{1, 2, 0, 4} {
		div60(value)
	}

	fmt.Println("hello")
}
