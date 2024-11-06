# Use Go as the base image
FROM golang:1.22.2

# Set the working directory
WORKDIR /app

# Copy the GoCelery worker application into the container
COPY . .

# Build the worker binary
RUN go build -o gocelery-worker worker.go

# Run the worker binary
CMD ["./gocelery-worker"]
