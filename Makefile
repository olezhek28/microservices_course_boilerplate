include .env

LOCAL_BIN:=$(CURDIR)/bin
LOCAL_MIGRATION_DIR=$(MIGRATION_DIR)
LOCAL_MIGRATION_DSN="host=localhost port=$(PG_PORT) dbname=$(PG_DBNAME) user=$(PG_USER) password=$(PG_PWD) sslmode=disable"

install-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3

lint:
	GOBIN=$(LOCAL_BIN) bin/golangci-lint run ./... --config .golangci.pipeline.yaml

install-deps:
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.4
	GOBIN=$(LOCAL_BIN) go install -mod=mod github.com/pressly/goose/v3/cmd/goose@v3.21.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod github.com/gojuno/minimock/v3/cmd/minimock@v3.3.6
	GOBIN=$(LOCAL_BIN) go install -mod=mod github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.21.0
	GOBIN=$(LOCAL_BIN) go install -mod=mod github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.21.0
	GOBIN=$(LOCAL_BIN) go install -mod=mod github.com/envoyproxy/protoc-gen-validate
	GOBIN=$(LOCAL_BIN) go install -mod=mod golang.org/x/perf/cmd/benchstat

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	mkdir -p pkg/user_v1
	protoc --proto_path api/user_v1 --proto_path vendor.protogen \
	--go_out=pkg/user_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/user_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	--grpc-gateway_out=pkg/user_v1 --grpc-gateway_opt=paths=source_relative \
	--plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
	--validate_out lang=go:pkg/user_v1 --validate_opt=paths=source_relative \
	--plugin=protoc-gen-validate=bin/protoc-gen-validate \
	--openapiv2_out=allow_merge=true,merge_file_name=api:api/user_v1/swagger \
	--plugin=protoc-gen-openapiv2=bin/protoc-gen-openapiv2 \
	api/user_v1/user.proto

add-vendor-protos:
	@if [ ! -d vendor.protogen/google ]; then \
  		mkdir  -p vendor.protogen/google &&\
  		git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
  		mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
  		rm -rf vendor.protogen/googleapis ;\
	fi
	@if [ ! -d vendor.protogen/validate ]; then \
      		git clone https://github.com/bufbuild/protoc-gen-validate vendor.protogen/validate.tmp &&\
      		mv vendor.protogen/validate.tmp/validate vendor.protogen &&\
      		rm -rf vendor.protogen/validate.tmp ;\
	fi
	@if [ ! -d vendor.protogen/protoc-gen-openapiv2 ]; then \
  			mkdir -p vendor.protogen/protoc-gen-openapiv2/options &&\
			git clone https://github.com/grpc-ecosystem/grpc-gateway vendor.protogen/openapiv2 &&\
			mv vendor.protogen/openapiv2/protoc-gen-openapiv2/options/*.proto vendor.protogen/protoc-gen-openapiv2/options &&\
			rm -rf vendor.protogen/openapiv2 ;\
	fi

build:
	GOOS=linux GOARCH=amd64 go build -o service_linux cmd/grpc_server/main.go

local-up:
	$(LOCAL_BIN)/goose -dir $(MIGRATION_DIR) postgres ${LOCAL_MIGRATION_DSN} up -v
local-down:
	$(LOCAL_BIN)/goose -dir $(MIGRATION_DIR) postgres ${LOCAL_MIGRATION_DSN} down -v

run:
	docker compose up -d

test-cover:
	go clean -testcache
	go test ./... -coverprofile=coverage.tmp.out -covermode count -coverpkg=\
github.com/neracastle/auth/internal/grpc-server/...,\
github.com/neracastle/auth/internal/usecases/...,\
github.com/neracastle/auth/internal/domain/...
	grep -v 'mocks\|config' coverage.tmp.out > coverage.out
	rm coverage.tmp.out
	go tool cover -html=coverage.out

test-cover-total:
	go tool cover -func=./coverage.out | grep "total" | awk '{print $$3}'