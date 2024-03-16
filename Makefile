BIN_NAME = Assembler
.PHONY: clean submit run build test source

all: build

build:
	@echo "Compiling Assembler using..."
	@go version
	@go build -o $(BIN_NAME) main.go
	@chmod +x Assembler

run:
	@./$(BIN_NAME) $(path)

submit: clean
	mkdir project06-jamesshima
	cp -r main.go Makefile asm LANGINFO test.sh tests project06-jamesshima
	zip -r project06-jamesshima.zip project06-jamesshima
	echo "ready for canvas"

clean:
	@rm -rf Assembler
	@rm -rf asm/*.hack

test:
	. ./test.sh