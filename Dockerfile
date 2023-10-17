# Use the alpine-based distribution of Go for a smaller final image.
FROM golang:1.17 AS builder


# Set the working directory inside the container.
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working directory inside the container.
COPY . .

# Build the application.

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

RUN chmod +x main


# Use a minimal alpine-based distribution for a smaller final image.
FROM golang:1.17-alpine

# Set the working directory inside the container.
WORKDIR /app

# Copy the binary from the builder stage.
COPY --from=builder --chmod=755 /app/main /app/main
RUN chmod +x /app/main

# Expose port 8080 to the outside world.
EXPOSE 8080

# Command to run the executable.
CMD ["/app/main"]
