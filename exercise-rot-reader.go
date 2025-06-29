package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(out []byte) (int, error) {

	in := make([]byte, len(out))
	n, err := rot.r.Read(in)

	if err != nil {
		return n, err
	}

	for i := 0; i < n; i++ {

		switch {
		case 65 <= in[i] && in[i] <= 90:
			out[i] = (((in[i] - 65) + 13) % 26) + 65
		case 97 <= in[i] && in[i] <= 122:
			out[i] = (((in[i] - 97) + 13) % 26) + 97
		default:
			out[i] = in[i]
		}
	}

	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
