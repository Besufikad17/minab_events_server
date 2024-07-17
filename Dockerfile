FROM golang:latest as builder

# Work directory
WORKDIR /minab_events

# Installing dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copying all the files
COPY . .

# Building the application
RUN go build -o minab_events

# Fetching the latest nginx image
FROM alpine:3.16 as production

# Certificates
RUN apk add --no-cache ca-certificates

# Copying built assets from builder
COPY --from=builder minab_events .

# Starting our application
CMD ./minab_events

# Exposing server port
EXPOSE 5000
