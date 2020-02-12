.PHONY:run
run:
	docker run -p 8080:8081 -it billybones256/grpcequation

.PHONY:grab
grab:
	docker pull billybones256/grpcequation:initail

.DEFAULT_GOAL := grab
