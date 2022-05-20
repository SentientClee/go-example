.PHONY: protobuf deps install

PROTOBUF_SERVICES_DIR := ./protobuf/services
PROTOBUF_SERVICES_GENERATED_OUT_REL_DIR := ../../generated/protobuf/services

protobuf:
	cd $(PROTOBUF_SERVICES_DIR) && \
	mkdir -p $(PROTOBUF_SERVICES_GENERATED_OUT_REL_DIR) && \
	for SERVICE_DIR in *; do \
		protoc \
			-I "." \
			-I "../../third_party/googleapis" \
			-I "$$SERVICE_DIR" \
			--go_out "$(PROTOBUF_SERVICES_GENERATED_OUT_REL_DIR)" --go_opt paths=source_relative \
			--go-grpc_out "$(PROTOBUF_SERVICES_GENERATED_OUT_REL_DIR)" --go-grpc_opt paths=source_relative \
			--grpc-gateway_out "$(PROTOBUF_SERVICES_GENERATED_OUT_REL_DIR)" \
			--grpc-gateway_opt paths=source_relative \
			$$SERVICE_DIR/*.proto; \
	done

deps:
	go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

install: deps protobuf
