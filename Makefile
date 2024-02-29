BIN_NAME = Assembler

build:
	@go build -o $(BIN_NAME) main.go

run:
	@./$(BIN_NAME) $(path)