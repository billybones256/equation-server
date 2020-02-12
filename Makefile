ifeq ($(OS),Windows_NT)
else
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		SUDO = 'sudo '
	endif
endif

.PHONY:proto
proto:
	protoc -I api/proto --go_out=plugins=micro:pkg/api api/proto/*.proto

.PHONY:build
build:
	$(SUDO) docker build -t equation-server .

.PHONY:push
push:
	$(SUDO) docker push billybones256/equation-server:initial

.PHONY:run
run:
	$(SUDO) docker run -it -p 50050:50051 -e MICRO_SERVER_ADDRESS=:50051 equation-server

.PHONY:grab
grab:
	$(SUDO) docker pull billybones256/equation-server:initial
