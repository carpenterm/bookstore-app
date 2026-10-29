# Build stage
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
COPY vendor/ ./vendor/
COPY . .

# Build main application
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o /app/main ./cmd/web

# Build seed binaries
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o /app/seed-gutenberg-books ./scripts/seed-gutenberg-books.go
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o /app/seed-images ./scripts/seed-images.go

# Final stage
FROM alpine:3.24
WORKDIR /app

# The build stage's public CA bundle, without a package download.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENV SSL_CERT_FILE=/etc/ssl/certs/ca-certificates.crt

# Copy main application
COPY --from=builder /app/main .
COPY templates ./templates
COPY static ./static

# Copy seed binaries
RUN mkdir -p /app/scripts/bin
COPY --from=builder /app/seed-gutenberg-books /app/scripts/bin/
COPY --from=builder /app/seed-images /app/scripts/bin/

# Copy bundled book covers (loaded by seed-images at init time)
COPY scripts/covers/ /app/covers/

# Copy migration files (for init jobs)
RUN mkdir -p /app/migrations
COPY migrations/*.sql /app/migrations/

EXPOSE 8080
CMD ["/app/main"]
