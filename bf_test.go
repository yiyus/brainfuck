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
		b, ok := <-bf.Out;
		if !ok {
			return;
		}
		fmt.Print(string(b));
	}
}

func TestAddition(t *testing.T) {
	prog := strings.Bytes(",>++++++[<-------->-],[<+>-]<.");
	bf := BrainFucker(prog, 30000);
	bi := "43";
	i := 0;
	for {
		oki := i<len(bi);
		if oki {
			bf.In <- bi[i];
			i++;
		}
		b, oko := <-bf.Out;
		if oko {
			fmt.Print(string(b));
		}
		if !oki && !oko {
			break;
		}
	}
}


