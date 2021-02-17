FROM golang:1.15

RUN apt-get update && \
    apt-get install -y protobuf-compiler && \
    rm -rf /var/lib/apt/lists/*
WORKDIR /app

RUN go mod init grpc && \
    go get github.com/golang/protobuf/protoc-gen-go@v1.4.3 && \
    go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.2.0  

ENTRYPOINT ["/usr/bin/protoc"]
CMD ["--help"]
