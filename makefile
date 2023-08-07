.PHONY: clean

BINARY_NAME=server

all: run

build:
	go build cmd/server/${BINARY_NAME}.go

run:
	./${BINARY_NAME}

start:
	make build
	make run

generate:
	go get github.com/99designs/gqlgen
	go run github.com/99designs/gqlgen generate

clean:
	go clean -cache
	rm -f ${BINARY_NAME}
dep:
	go mod download

lint:
	golangci-lint run --enable-all