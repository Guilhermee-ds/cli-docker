# Start from the official Golang image
FROM golang:1.22.7-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o app .

# Start a new stage from a smaller image
FROM alpine:3.18

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/app .

# Expose the application port
EXPOSE 8080

# Run the binary
CMD ["./app"]
