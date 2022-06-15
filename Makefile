
.PHONY: build
build:

	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o chat *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t gochat-charlie:latest