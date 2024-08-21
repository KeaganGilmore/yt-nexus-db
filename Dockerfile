# Copyright DWJ 2024.
# Distributed under the Boost Software License, Version 1.0.
# https://www.boost.org/LICENSE_1_0.txt

# Made with gpt-4o-mini

# Use the official Go image as a builder
FROM golang:1.22.4-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main.go

# Use a smaller base image to run the app
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the binary from the builder image
COPY --from=builder /app/main .

# Expose the port the app runs on
EXPOSE 8110

# Command to run the executable
CMD ["./main"]
