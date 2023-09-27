FROM golang

WORKDIR /app

COPY . .

EXPOSE 443

CMD go run cmd/main.go
