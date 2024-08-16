# Stage 1: Build the Go application
FROM golang:1.20-alpine AS builder

# Set environment variables
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Set the working directory
WORKDIR github.com/puneet105/job-queue

# Copy the Go module files
# COPY go.mod go.sum ./

# Download the Go modules
# RUN go mod download

# Copy the rest of the application code
COPY . .
RUN rm -rf go.* && \
    go mod init github.com/puneet105/job-queue && \
    go mod tidy 
# Build the Go application
RUN go build -o /github.com/puneet105/job-queue

# Stage 2: Create the final image
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /github.com/puneet105/job-queue .

# Copy the configuration file
COPY config.yaml .

# Expose the port the application will run on
EXPOSE 8080

# Command to run the application
CMD ["./job-queue"]
