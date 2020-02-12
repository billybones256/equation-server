FROM golang

WORKDIR /go/src/github.com/harry/grpcequation
COPY cmd/server cmd/server
COPY pkg pkg
WORKDIR /go/src/github.com/harry/grpcequation/cmd/server

RUN go get -d -v ./
RUN go build

CMD  ./server
EXPOSE 8081