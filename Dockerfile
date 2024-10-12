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

ENV DB_HOST=db
ENV DB_PORT=5432
ENV DB_USER=myuser
ENV DB_PASSWORD=mypassword
ENV DB_NAME=mydb
ENV DB_SSLMODE=disable

CMD ["./app"]
