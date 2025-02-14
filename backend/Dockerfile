FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o sylcot

EXPOSE 8080

CMD [ "./sylcot" ]