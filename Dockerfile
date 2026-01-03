FROM golang:1.24-bookworm

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080

CMD ["go", "run", "./cmd/crm/main.go"]

