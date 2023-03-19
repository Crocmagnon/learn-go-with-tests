FROM golang:1.20-alpine

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY acceptance_tests/ ./acceptance_tests/

ARG binToBuild

RUN go build -o svr acceptance_tests/cmd/${binToBuild}/*.go

EXPOSE 8080
CMD ["./svr"]