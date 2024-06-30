# Use the official Golang image to create a build artifact.
FROM golang:1.23rc1-alpine3.20 as builder

# Copy local code to the container image.
WORKDIR /app
COPY main.go .
COPY go.mod .

# Build the command inside the container.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o server

# Use a minimal image to run the server binary.
FROM alpine:3.20
COPY --from=builder /app/server /server

# Run the server on container startup.
CMD ["/server"]