FROM golang:latest

# Set the working directory to /eventpublisher
ADD . /go/src/eventpublisher
WORKDIR /go/src/eventpublisher

# Copy the Go application to the container
COPY main.go go.mod go.sum ./

# Build the Go application
RUN go get eventpublisher
RUN go install
RUN go build -o eventpublisher

# Expose the port the Go application will listen on
EXPOSE 8080

# Start the Go application
ENTRYPOINT ["/go/bin/eventpublisher"]