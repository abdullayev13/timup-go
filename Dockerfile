FROM golang

WORKDIR /app

COPY . .

ENV DOMAIN="16.16.182.36:443"

RUN go mod tidy

EXPOSE 443

CMD go run cmd/main.go
