FROM golang:1.24-alpine AS base

# Development stage
FROM base AS development
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/air-verse/air@latest
COPY . .
CMD ["air"]

# Build stage
FROM base AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -v -o ./app

# Production stage
FROM scratch AS prod
WORKDIR /prod
COPY --from=builder /build/app ./
COPY --from=builder /build/templates ./templates
COPY --from=builder /build/static ./static
ENV GIN_MODE=release
EXPOSE 8000
CMD ["./url-shortener"]
