# FROM golang:1.17 as builder
# WORKDIR /app
# COPY . ./
# RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server

# FROM alpine:3.14
# RUN apk add --no-cache ca-certificates
# COPY --from=builder /app/server /
# CMD ["/server"]


# Build Stage
FROM golang:1.17 AS builder
WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application files
COPY . ./

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server

# Runtime Stage
FROM alpine:3.14
RUN apk add --no-cache ca-certificates

# Copy the built binary from the builder stage
COPY --from=builder /app/server /server

# Command to run the application
CMD ["/server"]
