.PHONY:run
run:
	docker run -p 8080:8081 -it billybones256/grpcequation

.PHONY:grab
grab:
	docker pull billybones256/grpcequation:initial

.PHONY:client
client:
	go build -o client /cmd/client/main.go

.PHONY:client-run
client-run:
	./client ./cmd/client/test1.json

.DEFAULT_GOAL := grab
