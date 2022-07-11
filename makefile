.PHONY: client run-client

build:
	go build -o game.out

run-client: 
	go run main.go
