all: build

build:
	go build -o ./bin/app ./cmd/main.go

run-without-build:
	go run ./cmd

run:
	./bin/app

clean:
	rm -rf ./bin/*
