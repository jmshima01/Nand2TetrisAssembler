BIN_NAME = Assembler
.PHONY: clean submit run build test

all: build

build:
	@echo "Compiling Assembler using..."
	@go version
	@go build -o $(BIN_NAME) main.go
	@chmod +x Assembler

run:
	@./$(BIN_NAME) $(path)

submit:
	zip project06-jamesshima ./*
	echo "created zip for canvas"

test:
	@source test.sh

clean:
	@rm -rf Assembler
	@rm -rf *.hack