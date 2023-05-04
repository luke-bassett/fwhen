FROM golang:1.20-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go mod download

# Build the Go app
RUN go build -o main 

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]