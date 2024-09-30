include .env.local

LOCAL_BIN:=$(CURDIR)/bin

LOCAL_MIGRATION_DIR=$(MIGRATION_PATH)
LOCAL_MIGRATION_DSN="host=localhost port=$(POSTGRES_PORT) dbname=$(POSTGRES_DB) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) sslmode=disable"

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0
	GOBIN=$(LOCAL_BIN) go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.26
	GOBIN=$(LOCAL_BIN) go install golang.org/x/tools/cmd/goimports@latest
	GOBIN=$(LOCAL_BIN) go install github.com/gojuno/minimock/v3/cmd/minimock@latest
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u google.golang.org/grpc
	go get -u google.golang.org/grpc/reflection
	go get -u github.com/jackc/pgx/v5
	go get -u github.com/jackc/pgx/v5/pgxpool
	go get -u github.com/gojuno/minimock
	go get -u github.com/gojuno/minimock/v3
	go get -u github.com/brianvoe/gofakeit
	go get -u github.com/brianvoe/gofakeit/v6
	go get -u github.com/stretchr/testify/require
	go get -u github.com/joho/godotenv
	go get -u github.com/pkg/errors
	go get -u github.com/gomodule/redigo/redis
	go get -u google.golang.org/protobuf
	go get -u github.com/IBM/sarama

format:
	find . -name '*.go' -exec $(LOCAL_BIN)/goimports -w {} \;

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

goose-migration-status:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_PATH} postgres "host=$(POSTGRES_HOST) port=$(POSTGRES_PORT) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DB) sslmode=disable" status -v

goose-migration-up:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_PATH} postgres "host=$(POSTGRES_HOST) port=$(POSTGRES_PORT) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DB) sslmode=disable" up -v

goose-migration-down:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_PATH} postgres "host=$(POSTGRES_HOST) port=$(POSTGRES_PORT) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DB) sslmode=disable" down -v

generate:
	make generate-auth-api && \
	make generate-sqlc

generate-auth-api:
	mkdir -p pkg/user_v1
	protoc --proto_path api/user_v1 \
	--go_out=pkg/user_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/user_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/user_v1/user.proto

generate-sqlc:
	$(LOCAL_BIN)/sqlc generate
