FROM golang:1.23 as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd/ ./cmd
COPY pkg/ ./pkg

RUN go build -o /chat-app ./cmd

EXPOSE 8080

CMD ["/chat-app"]