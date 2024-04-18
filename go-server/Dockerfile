# Use the official Golang image to create a build artifact.
FROM golang:1.22 as builder

# Set the Current Working Directory inside the container.
WORKDIR /app

# Copy go.mod and go.sum if you are using Go modules.
COPY go.mod .
COPY go.sum .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed.
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container.
COPY . .

# Set the Current Working Directory inside the container to where your main.go is located.
WORKDIR /app

# Build the Go app
RUN GOOS=linux go build -o myapp .

FROM debian:bookworm-backports

WORKDIR /app
RUN apt-get update && apt-get install -y ncat && rm -rf /var/lib/apt/lists/*

# Copy the Pre-built binary file from the previous stage.
COPY --from=builder /app/myapp .

COPY docker.env .env
COPY wait-services.sh .
RUN chmod +x wait-services.sh

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/bin/bash", "./wait-services.sh", "./myapp"]
