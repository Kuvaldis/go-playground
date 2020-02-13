package main

import (
	"io"
	"os"
	"strings"
)

type Rot13Reader struct {
	r io.Reader
}

func (r Rot13Reader) Read(p []byte) (n int, err error) {
	rn, rerr := r.r.Read(p)
	for i := 0; i < rn; i++ {
		if isSmallLetter(p[i]) {
			if p[i] < 'n' {
				p[i] += 13
			} else {
				p[i] -= 13
			}
		} else if isBigLetter(p[i]) {
			if p[i] < 'N' {
				p[i] += 13
			} else {
				p[i] -= 13
			}
		}

	}
	return rn, rerr
}

func isSmallLetter(b byte) bool {
	return b >= 'a' && b <= 'z'
}

func isBigLetter(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

func main9() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := Rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
