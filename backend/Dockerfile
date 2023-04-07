FROM golang:1.19 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/app

FROM alpine:latest

WORKDIR /root

COPY --from=builder /app/main .
COPY --from=builder /app/configs/ /root/configs/

EXPOSE 8080

CMD ["./main"]