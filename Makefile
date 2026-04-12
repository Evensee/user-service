PROTOC = protoc
PROTO_DIR = protobuf-definitions
OUT_DIR = protobuf_generated


GO_BIN_DIR := $(shell go env GOBIN)
ifeq ($(strip $(GO_BIN_DIR)),)
GO_BIN_DIR := $(shell go env GOPATH)/bin
endif
PROTOC_GEN_GO = $(GO_BIN_DIR)/protoc-gen-go
PROTOC_GEN_GO_GRPC = $(GO_BIN_DIR)/protoc-gen-go-grpc

.PHONY: proto clean deps


deps:
	@echo "🔧 Installing protoc plugins..."
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


proto:
	@echo "🚀 Generating Go code from proto files..."
	@mkdir -p $(OUT_DIR)
	@$(PROTOC) \
		--plugin=protoc-gen-go=$(PROTOC_GEN_GO) \
		--plugin=protoc-gen-go-grpc=$(PROTOC_GEN_GO_GRPC) \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(OUT_DIR) \
		--go-grpc_out=$(OUT_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		$(shell find $(PROTO_DIR) -name "*.proto")


clean:
	@echo "🧹 Cleaning generated files..."
	@rm -rf $(OUT_DIR)


grpc:
	go run cmd/api/grpc/main.go

grpcup:
	docker compose up -d
	go run cmd/api/grpc/main.go

http:
	go run cmd/api/http/main.go

migrate:
	go run cmd/migrate/main.go

update_submodules:
	git submodule update --force --remote --init --recursive
