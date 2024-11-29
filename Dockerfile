FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY src/ ./

RUN go mod tidy
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
