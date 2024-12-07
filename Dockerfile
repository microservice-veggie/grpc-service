# Use the official Golang image to create a build artifact.
# This is the build stage.
FROM golang:1.23 AS builder

# Set the Current Working Directory inside the container
WORKDIR /grpc-service

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .
# Copy the .env file to the Working Directory inside the container
COPY .env .env

# Build the Go app for Linux
RUN GOOS=linux GOARCH=amd64 go build -o main .

# Start a new stage from scratch
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /grpc-service/main /root
COPY --from=builder /grpc-service/.env /root/
EXPOSE 8080
# Command to run the executable
ENTRYPOINT ["/root/main"]