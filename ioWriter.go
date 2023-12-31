package main

import (
	"fmt"
	"os"
	"io"
)

//Capper implements io.Writer and turns everything into uppercase
type Capper struct {
	wrt io.Writer
}

func (c *Capper) Write(p []byte) (n int, err error) {
	diff := byte('a' - 'A')

	out := make([]byte, len(p))
	for i,c := range p {
		if c >= 'a' && c <= 'z'{
			c -= diff
		}
		out[i] = c
	}

	return c.wrt.Write(out)
}

func main() {
	c := &Capper{os.Stdout}
	fmt.Fprintln(c, "Hello there")
}