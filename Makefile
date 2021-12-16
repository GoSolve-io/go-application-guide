USER = $(shell id -u)
GROUP = $(shell id -g)

.PHONY: build
build:
	go build ./cmd/app/...

.PHONY: run
run:
	go run ./cmd/app/...

# Run unit tests.
.PHONY: test
test:
	go test -race ./...

# For basic lint you can use:
# go vet ./... && golint ./...
# For more torough checks, we recommend golangci-lint with default configuration.
.PHONY: lint
lint: check-golangcilint-bin
	golangci-lint run --build-tags=integration --skip-dirs-use-default  --modules-download-mode=mod ./...

.PHONY: generate
generate: generate-proto generate-go

.PHONY: generate-proto
generate-proto:
	# Prepare docker image with generator.
	docker build -f ./api/Dockerfile -t apiprotoc ./api

	# Generate GRPC server files and openapi docs.
	docker run --rm -u ${USER}:${GROUP} \
		-v $(PWD):/app \
		apiprotoc \
		-I api/proto \
		--proto_path=/app/ \
		--go_out=plugins=grpc:. \
		--go_opt=module=github.com/nglogic/go-application-guide \
		--grpc-gateway_out=logtostderr=true:/app \
		--grpc-gateway_opt=module=github.com/nglogic/go-application-guide \
		--openapiv2_out=api/openapi \
		--openapiv2_opt=logtostderr=true \
		--openapiv2_opt=openapi_configuration=api/openapi/nglogic/bikerental/v1/service.swagger.config.yaml \
		api/proto/nglogic/bikerental/v1/service.proto

.PHONY: generate-go
generate-go: check-golangcilint-bin
	go generate ./...


### Checks ###

.PHONY: check-golangcilint-bin
check-golangcilint-bin:
ifeq (, $(shell which golangci-lint))
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.40.1
endif

check-mockgen:
ifeq (, $(shell which mockgen))
	go install github.com/golang/mock/mockgen@v1.5.0
endif
