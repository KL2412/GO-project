# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o urlshortener

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/urlshortener .
# Copy templates directory
COPY --from=builder /app/templates ./templates

# Expose the port
EXPOSE 8080

# Set environment variables with defaults
ENV BASE_URL=http://localhost:8080
ENV PORT=8080

# Run the application
CMD ["./urlshortener"] 