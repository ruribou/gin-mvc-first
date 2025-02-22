FROM golang:1.24-alpine

WORKDIR /app

RUN apk add --no-cache git

COPY src/go.mod src/go.sum ./
RUN go mod download

COPY src/ .

RUN go build -o /app/main ./cmd/main.go

EXPOSE 8080

CMD ["/app/main"]