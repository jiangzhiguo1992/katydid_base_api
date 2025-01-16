# deployments/docker/client/Dockerfile

# docker build -f deployments/docker/client/Dockerfile -t your-image-name .

# Use the official Golang image to create a build artifact.
FROM golang:1.23.4 AS builder
LABEL authors="jiang"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files from the root directory
#COPY ../../../go.mod ../../../go.sum ./
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
#COPY ../../../cmd ./cmd
COPY cmd ./cmd

# Build the Go app
RUN go build -o main ./cmd

# Start a new stage from scratch #ubuntu:24.04
FROM scratch

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main /app/main

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["/app/main"]