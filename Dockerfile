FROM golang:1.24-bookworm

WORKDIR /app

COPY . .

RUN go mod tidy

CMD ["make","up"]
CMD ["make","run"]
