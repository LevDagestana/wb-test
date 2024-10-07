FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM debian:buster-slim

COPY --from=builder /app/main /app/main

WORKDIR /app

CMD ["./main"]
