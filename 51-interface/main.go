package main

import (
	"fmt"
	"os"
)

func main() {
	fw := &FileWrite{Filename: "data.txt"}
	fmt.Fprintln(fw, "Hello World")

	//fmt.Fprintln(os.Stdout, "hello World")

	str := empty{}
	fmt.Fprintln(str, "Hello VW Folks")
}

type FileWrite struct {
	Filename string
}

type empty struct{}

func (f empty) Write(p []byte) (n int, err error) {
	print(string(p))
	return 0, nil
}

// io.Writer --> Write(p []byte) (n int, err error)
func (f *FileWrite) Write(p []byte) (n int, err error) {
	file, err := os.OpenFile(f.Filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	return file.Write(p)
}
