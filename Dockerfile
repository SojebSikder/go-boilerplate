FROM golang:1.24 AS prod

WORKDIR /app

COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o myapp ./cmd/main.go

EXPOSE 4000

CMD ["./myapp"]
