# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Ensure dependencies are in sync with the go.mod file
RUN go mod tidy

# Build the Go application
RUN go build -o anime-kuring .

# Expose the port on which the application will run
EXPOSE 8080

# Define the command to run the application when the container starts
CMD ["./anime-kuring"]
