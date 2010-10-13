include $(GOROOT)/src/Make.inc

TARG=brainfuck
GOFILES=brainfuck.go

include $(GOROOT)/src/Make.pkg

bf: main.go package
	$(GC) main.go && $(LD) -o $@ main.$O
