// Copyright 2009 JGL
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package brainfuck

import (
	"strings"
	"os"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	prog := `
		++++++++++[>+++++++>++++++++++>+++>+<<<<-]
		>++.>+.+++++++..+++.>++.<<+++++++++++++++.
		>.+++.------.--------.>+.>.!`
	bf := NewVM(prog, 30000, nil, os.Stdout)
	<-bf.Err
}

func TestAddition(t *testing.T) {
	prog := `,>++++++[<-------->-],[<+>-]<.`
	in := strings.NewReader("43")
	bf := NewVM(prog, 30000, in, os.Stdout)
	<-bf.Err
}
