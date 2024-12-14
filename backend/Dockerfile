FROM golang:1.23 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM debian:stable-slim

WORKDIR /app

COPY --from=builder /app/main .

RUN chmod +x ./main

EXPOSE 8383

CMD ["./main"]