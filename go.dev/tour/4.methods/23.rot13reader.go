package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13r rot13Reader) Read(b []byte) (int, error) {
	l, err := r13r.r.Read(b)
	if err != nil {
		return l, err
	}
	for i := range b {
		if 65 <= b[i] && b[i] <= 90 {
			if b[i] <= 77 {
				b[i] += 13
			} else {
				b[i] -= 13
			}
		} else if 97 <= b[i] && b[i] <= 122 {
			if b[i] <= 109 {
				b[i] += 13
			} else {
				b[i] -= 13
			}
		}
	}
	return l, nil
}

func reader_func2() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
