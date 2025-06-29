package main

import "fmt"

type MyReader struct{}

// Infinite reader
func (r MyReader) Read(buffer []byte) (int, error) {
	bufferLen := len(buffer)
	for i := 0; i < bufferLen; i++ {
		buffer[i] = 65
	}
	return bufferLen, nil
}

func main() {
	b := make([]byte, 8)
	r := MyReader{}
	fmt.Println(r.Read(b))
	fmt.Println(b)
}
