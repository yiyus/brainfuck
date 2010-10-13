// Copyright 2010 JGL
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Simple brainfuck interpreter

// Compile with:
// 8g main.go && 8l -o bf main.8

package main

import (
	"bitbucket.org/yiyus/brainfuck"
	"flag"
	"log"
	"io/ioutil"
	"os"
)

var size *int = flag.Int("s", 30000, "memory size")
var file *string = flag.String("f", "", "program file")

func main() {
	var prog string
	var err os.Error

	flag.Parse()
	log.SetFlags(0)
	if *file != "" {
		var f *os.File
		var p []uint8
		f, err = os.Open(*file, os.O_RDONLY, 0)
		if err != nil {
			log.Exitln("bf: ERROR", err)
		}
		p, err = ioutil.ReadAll(f)
		if err != nil {
			log.Exitln("bf: ERROR", err)
		}
		prog = string(p)
	} else {
		if flag.NArg() < 1 {
			log.Exitln("usage: bf [ -f program | program ]")
		}
		prog = flag.Arg(0)
	}
	bf := brainfuck.NewVM(prog, *size, os.Stdin, os.Stdout)
	if err = <-bf.Err; err != nil {
		log.Exitln("bf: ERROR", err)
	}
}
