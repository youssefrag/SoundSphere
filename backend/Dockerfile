# Use Golang base image with Go version 1.23 (or higher)
FROM golang:1.24-alpine

# Install dependencies and build tools
RUN apk update && apk add --no-cache git

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./

# Download Go dependencies
RUN go mod tidy
RUN go mod download

# Install Air (for hot-reloading)
RUN go install github.com/air-verse/air@latest

# Copy the rest of the application code
COPY . .

# Expose port
EXPOSE 8080

# Run the Go server with Air
CMD ["air"]
