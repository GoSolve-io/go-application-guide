.PHONY: build run lint proto

USER = $(shell id -u)
GROUP = $(shell id -g)

build:
	go build ./cmd/app/...

run:
	go run ./cmd/app/...

lint:
	go vet ./...
	golint ./...

proto:
	# Prepare docker image with generator.
	docker build -f ./api/Dockerfile -t apiprotoc ./api

	# Generate GRPC server files and openapi docs.
	docker run --rm -u ${USER}:${GROUP} \
		-v $(PWD):/app \
		apiprotoc \
		-I api \
		--proto_path=/app/ \
		--go_out=plugins=grpc:/app/internal/ \
		--grpc-gateway_out=logtostderr=true:/app/internal \
		--openapiv2_out api \
		--openapiv2_opt logtostderr=true \
		--openapiv2_opt openapi_configuration=api/service.swagger.config.yaml \
		api/service.proto
