PROTOC = protoc
PROTO_DIR = protobuf-definitions
OUT_DIR = protobuf_generated


PROTOC_GEN_GO = $(shell go env GOPATH)/bin/protoc-gen-go
PROTOC_GEN_GO_GRPC = $(shell go env GOPATH)/bin/protoc-gen-go-grpc

.PHONY: proto clean deps


deps:
	@echo "🔧 Installing protoc plugins..."
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


proto:
	@echo "🚀 Generating Go code from proto files..."
	@mkdir -p $(OUT_DIR)
	@$(PROTOC) \
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

http:
	go run cmd/api/http/main.go

migrate:
	go run cmd/migrate/main.go
