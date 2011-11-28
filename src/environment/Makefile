# include $(GOROOT)/src/Make.inc

# TARG=environment
# GOFILES=\
# 	cart_pole.go\
# 	environment.go

# include $(GOROOT)/src/Make.pkg


include $(GOROOT)/src/Make.inc
TARG=environment

all: _obj/$(TARG).a

GOFILES=\
	cart_pole.go\
	environment.go

_go_.6: $(GOFILES) core
	if [ ! -d _obj ]; then mkdir _obj; fi
	6g -I../core/_obj -o _go_.6 $(GOFILES)

_obj/environment.a: _go_.6 core
	gopack grc _obj/environment.a _go_.6

core:
	$(MAKE) -C ../core

clean:
	rm -rf _go_.6 _obj

.PHONY: core clean 
