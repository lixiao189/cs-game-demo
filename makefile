.PHONY: client run-client

client:
	cd ./client; \
	go build -o ../client.out \

run-client: client
	./client.out
