# Define the Docker image name and version
IMAGE_NAME=eventpublisherimage
IMAGE_VERSION=latest

# Define the Go application source code directory
APP_SRC_DIR=src/eventpublisher-app

# Define the Docker container name
CONTAINER_NAME=event-publisher

# Build the Docker image for the Go application
build:
	docker build -t $(IMAGE_NAME):$(IMAGE_VERSION) .

# Run the Go application inside a Docker container
run:
	docker run -d --name $(CONTAINER_NAME) -p 8080:8080 $(IMAGE_NAME):$(IMAGE_VERSION)

# Stop and remove the Docker container for the Go application
stop:
	docker stop $(CONTAINER_NAME)
	docker rm $(CONTAINER_NAME)