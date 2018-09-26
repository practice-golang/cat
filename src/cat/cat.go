package main

import (
	"fmt"
	"io"
	"os"
)

type textWriter struct{}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	tW := textWriter{}
	io.Copy(tW, f)
}

func (textWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("bytes 수:", len(bs))
	return len(bs), nil
}
