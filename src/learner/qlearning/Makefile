include $(GOROOT)/src/Make.inc

TARG=qlearning
GOFILES=\
	qlearning.go

default: _obj/$(TARG).a

_go_.6: $(GOFILES) core environment
	if [ ! -d _obj ]; then mkdir _obj; fi
	6g -I../../core/_obj -I../../environment/_obj -o _go_.6 $(GOFILES)

_obj/$(TARG).a: _go_.6 
	gopack grc _obj/$(TARG).a _go_.6

core:
	$(MAKE) -C ../../core

environment:
	$(MAKE) -C ../../environment

clean:
	rm -rf _go_.6 _obj

.PHONY: core environment clean 
