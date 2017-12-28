package main

import (
	"io"
	"os"

	"github.com/kuked/su8s"
)

func main() {
	w := su8s.Writer()
	_, err := io.Copy(w, os.Stdin)
	if err != nil {
		panic(err)
	}
}
