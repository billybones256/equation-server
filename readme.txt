protoc -I api/proto --go_out=plugins=grpc:pkg/api api/proto/solver.proto

evans api/proto/solver.proto -p 8081
