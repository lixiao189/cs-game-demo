.PHONY: client run-client run-server

build:
	go build -o game.out

run-client: 
	go run main.go

run-server:
	go run main.go -d
