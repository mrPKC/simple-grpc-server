# Go gRPC Server Example

This is a simple Go-based gRPC server example. It demonstrates how to create a basic gRPC server in Go, defining a service and implementing RPC methods.

## Prerequisites

- Go (1.17 or higher)
- Protobuf compiler (`protoc`)
- gRPC Go tools

## Quick Start

1. Clone the repository:

   ```bash
   git clone https://github.com/mrPKC/totality-corp-grpc.git
   cd your-grpc-server
   ```

2. Build and run the gRPC server:

   ```bash
   go run main.go
   ```

The gRPC server will be accessible on port 50051.

## Testing

You can test the gRPC server using gRPC clients or gRPC tools and on terminal you can use

```bash
go test
```

## Docker Support

You can also run the gRPC server in a Docker container. Use the provided `Dockerfile` and `Makefile` for image creation and container management.

- Build on Docker:

  ```bash
  make all
  ```