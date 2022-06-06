FROM golang:1.18 AS builder

WORKDIR /app

# Download necessary Go modules
COPY go.* ./

RUN go mod download

# COPY *.go ./
COPY *.go ./

RUN go build -o /github.com/akadendry/go-jdsport-raffle

# Make sure to expose the port the HTTP server is using
EXPOSE 8001
# Run the app binary when we run the container
ENTRYPOINT ["/github.com/akadendry/go-jdsport-raffle"]