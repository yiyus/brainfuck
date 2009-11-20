// Copyright 2009 JGL
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package brainfuck

import (
    "fmt";
    "strings";
    "testing";
)

func TestHelloWorld(t *testing.T) {
	prog := strings.Bytes("\
		++++++++++[>+++++++>++++++++++>+++>+<<<<-]\
		>++.>+.+++++++..+++.>++.<<+++++++++++++++.\
		>.+++.------.--------.>+.>.!");
	bf := BrainFucker(prog, 30000);
	for {
		b, ok := <- bf.out;
		if !ok {
			return;
		}
		fmt.Print(string(b));
	}
}
