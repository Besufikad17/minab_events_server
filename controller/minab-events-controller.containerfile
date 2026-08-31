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

# Build the Go app using the binary name expected by docker-compose.yml
RUN go build -o server .

# Expose port
EXPOSE 5000

CMD ["./server"]
