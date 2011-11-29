include $(GOROOT)/src/Make.inc

TARG=gorl

IPATH=-Icore/_obj -Ienvironment/_obj -Ilearner/_obj -Ilearner/qlearning/_obj
LPATH=-Lcore/_obj -Lenvironment/_obj -Llearner/_obj -Llearner/qlearning/_obj

all:$(TARG)

GOFILES=\
	gorlmain.go

$(TARG):gorlmain.6
	$(LD) $(LPATH) -o $@ $^

%.$O:%.go core environment learner qlearning
	$(GC) $(IPATH) -o $@ $<

core:
	$(MAKE) -C core

environment:
	$(MAKE) -C environment

learner:
	$(MAKE) -C learner

qlearning:
	$(MAKE) -C learner/qlearning

clean:
	rm -rf gorlmain.6
	(cd core && make clean)
	(cd environment && make clean)
	(cd learner && make clean)
	(cd learner/qlearning && make clean)

distclean: clean
	rm -f gorl

.PHONY:core environment learner qlearning clean

