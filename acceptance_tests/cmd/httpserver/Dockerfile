FROM golang:1.20-alpine

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY acceptance_tests/ ./acceptance_tests/

RUN go build -o svr acceptance_tests/cmd/httpserver/*.go

EXPOSE 8080
CMD ["./svr"]