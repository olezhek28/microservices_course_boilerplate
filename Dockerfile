FROM golang:1.20-alpine3.19 as builder
LABEL authors="ivansemeniv"

COPY . /neracastle/auth/src
WORKDIR /neracastle/auth/src

RUN go mod download
RUN go build -o ./bin/auth_server cmd/grpc-server/main.go

FROM alpine:3.19.2
WORKDIR /root/
COPY --from=builder /neracastle/auth/src/bin/auth_server .

CMD ["./auth_server"]