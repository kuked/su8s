package main

import (
	"io"
	"os"

	"github.com/kuked/su8s"
)

func main() {
	r := os.Stdin
	if len(os.Args) > 1 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		r = f
	}

	_, err := io.Copy(os.Stdout, su8s.Reader(r))
	if err != nil {
		panic(err)
	}
}
