# Use the official Golang image as the base image
FROM golang:1.19 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download and cache Go dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the appropriate binary based on the build argument
ARG TARGET
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/${TARGET}/${TARGET}.go

# Use a minimal Alpine image as the final base image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder image
COPY --from=builder /app/app .

# Start the binary based on the build argument
CMD ["./app"]
