package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (infA MyReader) Read(rb []byte) (n int, err error) {
	for n, err = 0, nil; n < len(rb); n++ {
		rb[n] = 'A'
	}
	return
}

func main() {
	reader.Validate(MyReader{})
}
