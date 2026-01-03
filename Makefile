.DEFAULT_GOAL := build

.PHONY:fmt vet build

fmt:
	go fmt ./...
vet: 
	go vet ./...
build: fmt vet
	go build
clean:
	go clean
run: fmt vet
	clear
	go run main