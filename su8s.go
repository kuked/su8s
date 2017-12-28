package su8s

import (
	"io"
	"os"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func Reader() io.Reader {
	d := japanese.ShiftJIS.NewDecoder()
	r := transform.NewReader(os.Stdin, d)

	return r
}

func Writer() io.Writer {
	e := japanese.ShiftJIS.NewEncoder()
	w := transform.NewWriter(os.Stdout, e)

	return w
}
