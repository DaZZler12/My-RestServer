# Use the official Go image as a base
FROM golang:1.17-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the entire current directory into the container
COPY . .

# Build the Go application
RUN go build -o My-RestServer .

# Start a new stage from scratch
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built executable from the previous stage
COPY --from=build /app/My-RestServer .

# Expose the port on which the Go application will listen
EXPOSE 8080

# Command to run the Go executable
CMD ["./My-RestServer"]
