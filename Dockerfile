# Use an official Go image as the base image
FROM golang:1.22.1 AS builder

# Set the working directory inside the container
WORKDIR /app

# Install protoc and its dependencies
RUN apt-get update && \
    apt-get install -y protobuf-compiler && \
    rm -rf /var/lib/apt/lists/*

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go modules
RUN go mod download

# Install protoc plugins for Go and gRPC
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN go install github.com/acoshift/goreload@latest

# Copy the rest of the application code
COPY . .

# Generate Go and gRPC code from proto files
RUN make generate

# Final image for running the application
FROM golang:1.22.1

# Set the working directory in the final image
WORKDIR /app

# Copy the generated code and other necessary files from the builder stage
COPY --from=builder /app .

# Expose the port on which your service runs (adjust based on your app)
EXPOSE 8082

