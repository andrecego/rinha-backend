from golang:1.21.0-alpine3.17 as builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build main.go

ENTRYPOINT ["/app/main"]
