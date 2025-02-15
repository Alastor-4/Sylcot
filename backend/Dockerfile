FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o sylcot

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/sylcot .

EXPOSE 8080

CMD ["./sylcot"]
