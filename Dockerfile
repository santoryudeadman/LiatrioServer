# syntax=docker/dockerfile:1

# Start from the official Go image
FROM golang:1.17-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download the Go dependencies
RUN go mod download

# Copy the source code into the container
COPY *.go ./

# Access the build arguments as environment variables


# Set environment variables, if required
# Build the Go application
RUN go build -o /LiatrioServer

# Expose the port your application listens on
EXPOSE 8080

# Run the Go application
CMD ["/LiatrioServer"]
