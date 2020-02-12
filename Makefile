.PHONY:proto
proto:
	protoc -I api/proto --go_out=plugins=micro:pkg/api api/proto/*.proto

.PHONY:build
build:
	docker build -t equation-server .

.PHONY:push
push:
	docker push billybones256/equation-server:initial

.PHONY:run
run:
	docker run -it -p 50050:50051 -e MICRO_SERVER_ADDRESS=:50051 equation-server

.PHONY:grab
grab:
	docker pull billybones256/equation-server:initial

