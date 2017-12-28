package main

import (
	"io"
	"os"

	"github.com/kuked/su8s"
)

func main() {
	r := su8s.Reader()
	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		panic(err)
	}
}
