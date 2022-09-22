proto:
	protoc -I pkg/pb/*.proto --go_out=plugins=grpc:.

server:
	go run cmd/main.go