// You won't believe how fucking proud I am of this piece of shit
// it took me like 3 fing hours,  because I was reading the wrong
// reader, so it was reading itself constantly and crashing.
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (self rot13Reader) Read(b []byte) (int, error) {
	buf := make([]byte, 8)
	i, e := self.r.Read(buf)
	if e != nil {
		return 0, e
	}

	for x, v := range buf {
		if x < i {
			if (v > 64 && v < 91) {
				t := v - 13
				if t < 65 {
					b[x] = t + 26
					continue
				}
				b[x] = t
				continue
			}
			if (v > 96 && v < 123) {
				t := v - 13
				if t < 97 {
					b[x] = t + 26
					continue
				}
				b[x] = t
				continue
			}
			b[x] = v
			continue
		}
		b[x] = 0
	}
	return 8, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}