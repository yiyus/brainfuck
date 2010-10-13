include $(GOROOT)/src/Make.inc

TARG=brainfuck
GOFILES=bf.go

include $(GOROOT)/src/Make.pkg

bf: main.go package
	$(GC) main.go && $(LD) -o $@ main.$O
