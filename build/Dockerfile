# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/app -ldflags="-s -w" main.go

# Runtime stage
FROM scratch

WORKDIR /app
COPY --from=builder /app/app /app/app

EXPOSE 3000
CMD ["/app/app", "api"]