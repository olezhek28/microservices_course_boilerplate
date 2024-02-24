FROM golang:1.20.3-alpine AS builder

COPY . /github.com/dmtrybogdanov/auth/source
WORKDIR /github.com/dmtrybogdanov/auth/source

RUN go mod download
RUN go build -o ./bin/auth cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/dmtrybogdanov/auth/source/bin/ .

CMD ["./auth"]