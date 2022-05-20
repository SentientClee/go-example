REGISTRY := TODO
IMAGE_TAG := $(shell git rev-parse --short HEAD)
IMAGE_NAME := example:$(IMAGE_TAG)

go-build:
	GOOS=linux GOARCH=amd64 go build -o example

build: go-build
	docker build -t $(REGISTRY)/$(IMAGE_NAME) .

push:
	docker push $(REGISTRY)/$(IMAGE_NAME)
