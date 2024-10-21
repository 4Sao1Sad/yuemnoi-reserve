# Makefile for generating Go and gRPC code from protobuf files

PROTOC_CMD = protoc
PROTO_DIR = proto # The directory containing your .proto files
GO_OUT = --go_out=.
GO_OPT = --go_opt=paths=source_relative
GRPC_OUT = --go-grpc_out=.
GRPC_OPT = --go-grpc_opt=paths=source_relative
PROTO_FILES = $(shell find $(PROTO_DIR) -name '*.proto')

.PHONY: all generate clean

tools-install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	
all: generate dev

# Generates the Go and gRPC code from proto files
# The --go_opt=paths=source_relative ensures files are placed next to the proto files
generate:
	$(PROTOC_CMD) $(GO_OUT) $(GO_OPT) $(GRPC_OUT) $(GRPC_OPT) $(PROTO_FILES)

# Cleans up generated files
clean:
	@echo "Cleaning generated files..."
	find $(PROTO_DIR) -name '*.pb.go' -delete

dev: 
	goreload --build=cmd/