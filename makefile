.PHONY: client run-client run-client1 run-server

build:
	go build -o game.out

run-client: 
	go run main.go -n "node"

run-client1:
	go run main.go -n "s1ain"

run-server:
	go run main.go -d
