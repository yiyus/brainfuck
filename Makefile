include $(GOROOT)/src/Make.inc

TARG=brainfuck
GOFILES=brainfuck.go

CLEANFILES+=bf

include $(GOROOT)/src/Make.pkg

main.$O: main.go package
	$(GC) -I_obj $<

bf: main.$O
	$(LD) -L_obj -o $@ $<
