FROM golang:1.22.2-alpine3.19

WORKDIR /app

COPY . .

RUN go mod download

CMD ["go", "run", "main.go"]

EXPOSE 8080
