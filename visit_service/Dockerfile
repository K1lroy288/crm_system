FROM golang:alpine AS builder

WORKDIR /visit_service

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /visit_service

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /visit_service/visit-service .

COPY .env .

EXPOSE 8001

CMD ["./visit-service"]
