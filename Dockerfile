# Use a minimal base image for Go apps
FROM images.mahdifarahmand.ir/golang:1.25.4-alpine3.21

# Set the working directory inside the container
WORKDIR /app

# Copy the source code
COPY . .

# Download dependencies
RUN go mod download

# Build the app
RUN go build -o server

# Expose the port your app uses
EXPOSE 8080

# Command to run the binary
CMD ["./server"]
