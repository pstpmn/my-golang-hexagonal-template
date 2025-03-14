# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install necessary dependencies including tzdata
RUN apk add --no-cache git tzdata

# Set timezone to Asia/Bangkok
ENV TZ=Asia/Bangkok

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/app -ldflags="-s -w" main.go

# Runtime stage
FROM scratch

WORKDIR /app

# Copy timezone data
COPY --from=builder /usr/share/zoneinfo/Asia/Bangkok /usr/share/zoneinfo/Asia/Bangkok
ENV TZ=Asia/Bangkok

COPY --from=builder /app/app /app/app

EXPOSE 3000
CMD ["/app/app", "api"]
