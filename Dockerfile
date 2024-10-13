# Stage 1: Builder
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o app ./cmd/main.go || { echo 'Build failed'; exit 1; }

RUN ls -la /app

RUN chmod +x /app/app

FROM alpine:latest

RUN apk add --no-cache postgresql-client

WORKDIR /root/

COPY --from=builder /app/app .

ARG DB_HOST
ARG DB_PORT
ARG DB_USER
ARG DB_PASSWORD
ARG DB_NAME
ARG DB_SSLMODE

ENV DB_HOST=${DB_HOST}
ENV DB_PORT=${DB_PORT}
ENV DB_USER=${DB_USER}
ENV DB_PASSWORD=${DB_PASSWORD}
ENV DB_NAME=${DB_NAME}
ENV DB_SSLMODE=${DB_SSLMODE}

CMD ["./app"]
