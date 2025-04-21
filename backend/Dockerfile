# Step 1: Use an official Golang image as the base image
FROM golang:1.24-alpine

# Set the working directory inside the container to the root of the backend
WORKDIR /backend

# Copy the Go modules files (go.mod and go.sum)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application files into the container
COPY . .

# Expose the port the Go application will run on
EXPOSE 8080

# Define the command to run the Go application
CMD ["go", "run", "main.go"]
