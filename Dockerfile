FROM golang:latest

ENV GO111MODULE=on
ENV PORT=8080

WORKDIR /app

# Copy go.mod and go.sum first for optimized caching
COPY go.mod go.sum ./
RUN go mod download

# Now copy the whole source code
COPY . .

# Build the application
RUN go build -o main .

# Change the permission of the application
RUN chmod +x main
