package main

import (
	"bytes"
	"io"
	"log"
	"os"
)

type QuoteToTickReader struct {
	io.Reader
}

func (r QuoteToTickReader) Read(p []byte) (n int, err error) {
	ticked := make([]byte, len(p))
	n, err = r.Reader.Read(ticked)
	ticked = ticked[:n]
	for r, c := range map[rune]byte{
		'“': '"', '”': '"',
		'‘': '\'', '’': '\'',
	} {
		ticked = bytes.ReplaceAll(
			ticked,
			[]byte(string(r)),
			[]byte{c},
		)
	}
	copy(p, ticked)
	return n, err
}

func main() {
	var r io.Reader = os.Stdin
	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		r = file
	}
	r = &QuoteToTickReader{r}
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}
