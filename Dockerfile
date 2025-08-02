FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod ./go.mod
COPY go.sum ./go.sum
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /any-api ./cmd/api

EXPOSE 8080

CMD ["/any-api"]