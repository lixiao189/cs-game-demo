.PHONY: client run-client

client:
	go build -o client.out \

run-client: client
	./client.out
