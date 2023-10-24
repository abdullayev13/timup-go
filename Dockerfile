FROM golang

WORKDIR /app

COPY . .

ENV DOMAIN="16.16.182.36:443"
ENV PORT=443
#   postgres db
ENV DB_HOST='localhost'
ENV DB_PORT=5432
ENV DB_DATABASE='postgres'
ENV DB_USERNAME='postgres'
ENV DB_PASSWORD='password'

RUN go mod tidy

EXPOSE 443

CMD go run cmd/main.go
