FROM golang

WORKDIR /app

COPY . .

RUN go mod tidy

EXPOSE 443

CMD go run cmd/main.go
