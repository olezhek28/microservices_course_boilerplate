FROM golang:1.21.1-alpine as builder
COPY . /github.com/MikhailRibalkov/chat-server/pkg/chatServer_v1/source/
WORKDIR /github.com/MikhailRibalkov/chat-server/pkg/chatServer_v1/source/

RUN go mod download
RUN go build -o ./bin/chat_server cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/MikhailRibalkov/chat-server/pkg/chatServer_v1/source/bin/chat_server .

CMD ["./chat_server"]