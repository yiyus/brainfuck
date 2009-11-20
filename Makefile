include $(GOROOT)/src/Make.$(GOARCH)

TARG=brainfuck
GOFILES=bf.go

include $(GOROOT)/src/Make.pkg
