# syntax=docker/dockerfile:1

# FROM golang:1.19

# # Set destination for COPY
# WORKDIR /app

# # Download Go modules
# COPY go.mod go.sum ./
# RUN go mod download

# # Copy the source code. Note the slash at the end, as explained in
# # https://docs.docker.com/engine/reference/builder/#copy
# COPY . .

# # Build
# RUN go build -o /app-build

# # Optional:
# # To bind to a TCP port, runtime parameters must be supplied to the docker command.
# # But we can document in the Dockerfile what ports
# # the application is going to listen on by default.
# # https://docs.docker.com/engine/reference/builder/#expose
# EXPOSE 3000

# # Run
# CMD ["/app-build"]

# Start from a minimal golang base image
FROM golang:alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Start a new stage from scratch
FROM alpine:latest

# Set necessary environment variables
# ENV PORT=3000

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/app /app/app

# Expose port 3000 to the outside world
EXPOSE 3000