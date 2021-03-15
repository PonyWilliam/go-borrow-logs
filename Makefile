GOPATH:=$(shell go env GOPATH)
.PHONY: proto
proto:
    
	protoc --proto_path=. --micro_out=. --go_out=:. proto/borrowLogs.proto
    

.PHONY: build
build: proto

	go build -o borrowLogs-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t borrowLogs-service:latest