# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest
ENV GO111MODULE=auto

# Add Maintainer Info
LABEL maintainer="Redes2_G17"

# Set the Current Working Directory inside the container
WORKDIR /

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 5000
EXPOSE 3000

# Command to run the executable
CMD ["./main"]