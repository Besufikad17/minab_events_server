# Use the official Golang image as the base image
FROM golang:1.22.1

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build the Go app with server.go as the entry point
RUN go build -o main server.go

# Expose port 8080 to the outside world
EXPOSE 5000
