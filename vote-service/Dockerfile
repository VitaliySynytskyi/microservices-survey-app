# Use the official Golang image as a base image
FROM golang:1.22.3-alpine

# Create a directory for the application
RUN mkdir /app

# Add the application code to the /app directory
ADD . /app

# Set the working directory to /app
WORKDIR /app

# Build the Go application and name the output binary "main"
RUN go build -o main .

# Specify the command to run the application
CMD ["/app/main"]
