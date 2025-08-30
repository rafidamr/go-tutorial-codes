package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (int, error) {
	b[0] = 65
	return 1, nil
}

func reader_func() {
	reader.Validate(MyReader{})
}
