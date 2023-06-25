
build:
	go build cmd/main.go
run:
	go run cmd/main.go
proto:
	protoc --go_out=plugins=grpc:. protos/catalog.proto