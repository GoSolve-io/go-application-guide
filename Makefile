.PHONY: build run test lint proto check-golangcilint-bin

USER = $(shell id -u)
GROUP = $(shell id -g)

build:
	go build ./cmd/app/...

run:
	go run ./cmd/app/...

test:
	go test -race ./cmd/app/...

# For basic lint you can use:
# go vet ./... && golint ./...
# For more torough checks, we recommend golangci-lint with default configuration.
lint: check-golangcilint-bin
	golangci-lint run ./...

proto:
	# Prepare docker image with generator.
	docker build -f ./api/Dockerfile -t apiprotoc ./api

	# Generate GRPC server files and openapi docs.
	docker run --rm -u ${USER}:${GROUP} \
		-v $(PWD):/app \
		apiprotoc \
		-I api/proto \
		--proto_path=/app/ \
		--go_out=plugins=grpc:/app/api/ \
		--grpc-gateway_out=logtostderr=true:/app/api \
		--openapiv2_out api/openapi \
		--openapiv2_opt logtostderr=true \
		--openapiv2_opt openapi_configuration=api/openapi/nglogic/bikerental/v1/service.swagger.config.yaml \
		api/proto/nglogic/bikerental/v1/service.proto

	# Move generated package to proper path.
	rm -rf ./pkg/api/*
	cp -rf ./api/github.com/nglogic/go-example-project/pkg/api/* ./pkg/api/
	rm -rf ./api/github.com


check-golangcilint-bin:
ifeq (, $(shell which golangci-lint))
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.37.1
endif
