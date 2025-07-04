FROM golang:1.24-alpine AS base

# Development stage
FROM base AS dev
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/air-verse/air@latest
COPY . .
EXPOSE 80
ENTRYPOINT ["air"]

# Build stage
FROM base AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -v -o ./app
RUN go build ./tools/cli.go

# Production stage
FROM scratch AS prod
WORKDIR /prod
COPY --from=builder /build/app ./app
COPY --from=builder /build/cli ./cli
COPY --from=builder /build/templates ./templates
COPY --from=builder /build/static ./static
EXPOSE 80
ENTRYPOINT ["./app"]
