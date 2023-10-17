# Use the official Go image as a parent image
FROM golang:1.17

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. If they exist in your project.
RUN go mod download

# Copy the source from the current directory to the working directory inside the container
COPY . .

# Build the application
RUN go build -o main .

# Expose port 8080 for the application
EXPOSE 8080

# Command to run the application
CMD ["./main"]
