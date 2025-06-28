FROM golang:1.24-alpine AS base

# Development stage
FROM base AS development
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/air-verse/air@latest
COPY . .
CMD ["air"]
