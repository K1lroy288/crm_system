FROM golang:alpine AS builder

WORKDIR /user_service

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /user_service

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /user_service/user-service .

COPY .env .

EXPOSE 8000

CMD ["./user-service"]
