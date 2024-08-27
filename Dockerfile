# Use the official Golang image as the base
FROM golang:1.23

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application code (main.go) into the container
COPY main.go .

# Initialize Go module, download dependencies, and build the application
RUN go mod init myapp && \
    go mod tidy && \
    go mod download

# Expose port 8000 for the application
EXPOSE 8000

# Command to run the Go application -> Entry point
CMD ["go", "run", "main.go"]