# Use multi-stage for smaller final image
FROM golang:1.24 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o myapp ./main.go

# Final minimal image
FROM debian:bullseye-slim
WORKDIR /app

COPY --from=builder /app/myapp .

# Port your Gin server uses (e.g., 4000)
EXPOSE 4000

CMD ["./myapp"]
