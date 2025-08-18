# Step 1: Build stage
FROM golang:1.24.6-alpine3.22 AS builder
WORKDIR /app

# Copy only go.mod first (for caching dependencies)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the Go app (assuming main.go is inside /cmd/)
RUN go build -o myapp ./cmd

# Step 2: Run stage
FROM alpine:3.22
WORKDIR /root/

# Copy only the binary (smaller image)
COPY --from=builder /app/myapp .

EXPOSE 8080
ENTRYPOINT ["./myapp"]