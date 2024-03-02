BIN_NAME = Assembler

build:
	@go build -o $(BIN_NAME) main.go

run:
	@./$(BIN_NAME) $(path)

submit:
	zip project06-jamesshima ./*
	echo "created zip for canvas"