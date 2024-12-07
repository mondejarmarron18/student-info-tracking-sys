FROM golang:1.23.3-alpine3.20

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.* .

RUN go mod download

COPY . .

CMD ["air", "-c", ".air.toml"]