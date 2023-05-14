module = $(shell go list -m)

main = main

image = main

default:
	@echo "Usage: \n  make [command]\n"
	@echo "Available Commands: \n  gen-grpc\n  build-bin main=entry-name\n  img image=image-name"

gen-grpc:
	protoc --go_out=./ --go_opt=module=${module} --go-grpc_out=./ --go-grpc_opt=module=${module} --grpc-gateway_out=:./ --grpc-gateway_opt=logtostderr=true --grpc-gateway_opt=module=${module} proto/*.proto

build-bin:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/${main} cmd/${main}.go

img:
	docker build -f build/Dockerfile -t ${image}:latest .

