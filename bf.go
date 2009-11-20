// Copyright 2009 JGL
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
	Brainfuck virtual machines.

	Commands:
		> 	increment the data pointer (to point to the next cell to the right).
		< 	decrement the data pointer (to point to the next cell to the left).
		+ 	increment (increase by one) the byte at the data pointer.
		- 	decrement (decrease by one) the byte at the data pointer.
		. 	output the value of the byte at the data pointer.
		, 	accept one byte of input, storing its value in the byte at the data pointer.
		[ 	if the byte at the data pointer is zero, then instead of moving the instruction
			pointer forward to the next command, jump it forward to the command after
			the matching ] command*.
		] 	if the byte at the data pointer is nonzero, then instead of moving the instruction
			pointer forward to the next command, jump it back to the command after the
			matching [ command*.
*/

package brainfuck

// BrainFuckVM represents the input and output of the virtual machine
type BrainFuckVM struct {
	in	chan byte;
	out	chan byte;
}

func (bf BrainFuckVM) core(prog []byte, size int) {
	a := make([]byte, size);
	if len(prog) == 0 || size == 0 {
		return
	}
	p := 0;
	pc := 0;
	for {
		switch prog[pc] {
		case '>':
			p++
		case '<':
			p--
		case '+':
			a[p]++
		case '-':
			a[p]--
		case '.':
			bf.out <- a[p]
		case ',':
			// test/turing.go cannot do this!
			a[p] = <-bf.in
		case '[':
			if a[p] == 0 {
				for nest := 1; nest > 0; pc++ {
					switch prog[pc+1] {
					case ']':
						nest--
					case '[':
						nest++
					}
				}
			}
		case ']':
			if a[p] != 0 {
				for nest := -1; nest < 0; pc-- {
					switch prog[pc-1] {
					case ']':
						nest--
					case '[':
						nest++
					}
				}
			}
		}
		pc++;
		if pc == len(prog) {
			return
		}
	}
}

// BrainFucker launchs a new virtual machine with the specified
// program and memory and return a BrainFuckVM struct to
// access to its I/O ports
func BrainFucker(prog []byte, size int) *BrainFuckVM {
	bf := new(BrainFuckVM);
	bf.in = make(chan byte);
	bf.out = make(chan byte);
	go bf.core(prog, size);
	return bf;
}
