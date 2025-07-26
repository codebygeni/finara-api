# Start from the official Go base image
FROM golang:1.23.0 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .
RUN test -f config/firebase-service-account.json || echo "⚠️ Missing Firebase key! Cloud Run may fail if this is not mounted or baked in."

# Build the Go app
RUN go build -o geni-firestore-api main.go

# Start a new minimal base image
FROM gcr.io/distroless/base-debian11

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary and config
COPY --from=builder /app/geni-firestore-api .
COPY --from=builder /app/config /app/config


# Expose the port Cloud Run expects
EXPOSE 9090

# Set the entry point
ENTRYPOINT ["./geni-firestore-api"]