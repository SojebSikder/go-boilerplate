# -------- Builder Stage --------
FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o myapp ./cmd/main.go

# -------- Final Stage --------
FROM gcr.io/distroless/base-debian12:nonroot

WORKDIR /app

COPY --from=builder /app/myapp .

EXPOSE 4000

ENTRYPOINT [ "./myapp" ]
CMD ["server"]
