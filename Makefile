SRC := $(shell find . -type f -name '*.go' -print) go.mod go.sum

bin/todo: $(SRC)
	go build -o bin/todo ./cmd/todo
