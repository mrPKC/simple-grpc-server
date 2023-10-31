.PHONY: all build run clean

# Set the name of your Docker image and container
DOCKER_IMAGE_NAME = mygrpcapp-image
DOCKER_CONTAINER_NAME = mygrpcapp-container

# Define the port mappings for your gRPC server
DOCKER_GRPC_PORT = 50051
HOST_GRPC_PORT = 50051

all: build run

build:
	docker system prune -a --volumes
	docker build -t $(DOCKER_IMAGE_NAME) .

run:
	docker run -p $(HOST_GRPC_PORT):$(DOCKER_GRPC_PORT) --name $(DOCKER_CONTAINER_NAME) $(DOCKER_IMAGE_NAME)

clean:
	docker stop $(DOCKER_CONTAINER_NAME)
	docker rm $(DOCKER_CONTAINER_NAME)
	docker rmi $(DOCKER_IMAGE_NAME)
