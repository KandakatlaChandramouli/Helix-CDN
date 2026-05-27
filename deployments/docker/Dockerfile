FROM golang:1.25

WORKDIR /app

COPY . .

RUN go build -o helix-edge cmd/edge/main.go

CMD ["./helix-edge"]
