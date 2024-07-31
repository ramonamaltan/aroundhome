# Use the official Golang image to create a build artifact
FROM golang:1.18 AS builder
# Set the Current Working Directory inside the container
WORKDIR /app
# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
# Copy the source code into the container
COPY cmd/ cmd/
COPY internal/api/ internal/api/
COPY internal/db/ internal/db/
COPY internal/models/ internal/models/
# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api-server ./cmd

# Start a new stage from scratch
FROM alpine:latest  
# Install necessary packages for Alpine
RUN apk add --no-cache ca-certificates curl
# Set the Current Working Directory inside the container
WORKDIR /root/
# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/api-server .
# Expose port 8080 to the outside world
EXPOSE 8080
# Command to run the executable
CMD ["./api-server"]
