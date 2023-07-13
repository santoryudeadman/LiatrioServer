# Start from the official Go image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download the Go dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Access the build arguments as environment variables
ARG API_KEY
ARG SERVER_CERT
ARG SERVER_KEY

# Set environment variables, if required
ENV API_KEY=$API_KEY
ENV SERVER_CERT=$SERVER_CERT
ENV SERVER_KEY=$SERVER_KEY

# Build the Go application
RUN go build -o app

# Expose the port your application listens on
EXPOSE 8080

# Run the Go application
CMD ["./app"]
