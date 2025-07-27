FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod ./go.mod
COPY go.sum ./go.sum
RUN go mod download

COPY . .

RUN go build -o /any-api ./cmd/api

EXPOSE 8080

CMD ["/any-api"]
