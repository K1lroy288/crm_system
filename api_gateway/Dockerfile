FROM golang:alpine AS builder

WORKDIR /api_gateway

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /api_gateway

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /api_gateway/api-gateway .

COPY --from=builder /api_gateway/templates ./templates/

COPY .env .

EXPOSE 8000

CMD ["./api-gateway"]
