FROM golang:latest AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files for caching dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api/*.go

FROM scratch AS prod

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Command to run the executable
CMD ["/main"]