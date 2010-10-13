// Copyright 2009 JGL
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package brainfuck

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	prog := []byte(`
		++++++++++[>+++++++>++++++++++>+++>+<<<<-]
		>++.>+.+++++++..+++.>++.<<+++++++++++++++.
		>.+++.------.--------.>+.>.!`)
	bf := NewVM(prog, 30000)
	for {
		b, ok := <-bf.Out
		if !ok {
			return
		}
		fmt.Print(string(b))
	}
}

func TestAddition(t *testing.T) {
	prog := []byte(`,>++++++[<-------->-],[<+>-]<.`)
	bf := NewVM(prog, 30000)
	bi := "43"
	i := 0
	for {
		reading := i < len(bi)
		if reading {
			bf.In <- bi[i]
			i++
		}
		b, writing := <-bf.Out
		if writing {
			fmt.Print(string(b))
		}
		if !reading && !writing {
			break
		}
	}
}
