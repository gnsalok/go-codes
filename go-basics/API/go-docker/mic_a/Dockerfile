# Use the official Golang image as a base image
FROM golang:1.17-alpine AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go modules files
# COPY go.mod go.sum ./
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o mic_a .

# Use a minimal Alpine image to reduce the image size
FROM alpine:latest

# Copy the executable built in the previous stage
COPY --from=builder /app/mic_a /app/mic_a

# Run the microservice
CMD ["/app/mic_a"]
