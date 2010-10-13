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

	See also: http://golang.org/test/turing.go (many thanks to the author of this code!)
*/

package brainfuck

import(
	"io"
	"os"
)

// A VM is a brainfuck virtual machine.
// The Err chanel can be read once the program
// finishes. A nil value means the program run
// sucessfully, other os.Error value is an io error.
type VM struct {
	Err	chan os.Error
	in	io.Reader
	out	io.Writer
	size	int
}

func (vm VM) run(prog string) {
	a := make([]byte, vm.size)
	if len(prog) == 0 || vm.size == 0 {
		return
	}
	p := 0
	pc := 0
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
			if _, err := vm.out.Write(a[p:p + 1]); err != nil {
				vm.Err <- err
			}
		case ',':
			// test/turing.go cannot do this!
			if _, err := vm.in.Read(a[p:p + 1]); err != nil {
				vm.Err <- err
			}
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
		pc++
		if pc == len(prog) {
			vm.Err <- nil
			return
		}
	}
}

// NewVM returns a new virtual machine running the specified program,
// with the given memory size, reading from in and writting to out.
func NewVM(prog string, size int, in io.Reader, out io.Writer) *VM {
	vm := &VM{make(chan os.Error), in, out, size}
	go vm.run(prog)
	return vm
}
