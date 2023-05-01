FROM golang:latest

# Set the working directory to /eventpublisher-app
WORKDIR /eventpublisher-app

# Copy the Go application to the container
COPY main.go go.mod go.sum ./

# Build the Go application
RUN go build -o eventpublisher

# Expose the port the Go application will listen on
EXPOSE 8080

# Start the Go application
CMD ["/eventpublisher-app/eventpublisher"]