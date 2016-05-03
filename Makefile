
all: build

build:
	go build

test: build
	./hash-file -o testdata/out1.out testdata/in1.txt testdata/in2.txt
	echo diff testdata/out1.out testdata/out1.ref

install: build
	cp hash-file ~/bin

