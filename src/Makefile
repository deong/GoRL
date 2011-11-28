include $(GOROOT)/src/Make.inc

TARG=gorl

all:$(TARG)

GOFILES=\
	gorlmain.go

$(TARG):gorlmain.6
	$(LD) -Lcore/_obj -Lenvironment/_obj -o $@ $^

%.$O:%.go core environment
	$(GC) -Icore/_obj -Ienvironment/_obj -o $@ $<

core:
	$(MAKE) -C core

environment:
	$(MAKE) -C environment

clean:
	rm -rf gorlmain.6
	(cd core && make clean)
	(cd environment && make clean)

distclean: clean
	rm -f gorl

.PHONY:core environment clean

