# Stage 1: Build the Go application
FROM golang:1.23-alpine AS builder
WORKDIR /app
RUN apk add --no-cache build-base
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 go build -o main .
RUN ls -la main
FROM alpine:latest
WORKDIR /app
RUN apk add --no-cache sqlite-dev
COPY --from=builder /app/main .
COPY app/app.db /app/app.db
RUN chmod +x /app/main
EXPOSE 3000
CMD ["/app/main"]