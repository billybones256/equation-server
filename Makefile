.PHONY:proto
proto:
	protoc -I api/proto --go_out=plugins=micro:pkg/api api/proto/*.proto

.PHONY:push
push:
	docker push billybones256/equation-client:initial

.PHONY:build
build:
	docker build -t equation-client .

.PHONY:run
run:
	docker run -it -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 equation-client

.PHONY:grab
grab:
	docker pull billybones256/equation-client:initial
