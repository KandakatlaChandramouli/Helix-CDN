FROM golang:1.25 AS builder

WORKDIR /app

COPY . .

RUN go build -o helix-edge ./cmd/edge
RUN go build -o helix-reactor ./cmd/reactor

FROM debian:bookworm-slim

WORKDIR /runtime

COPY --from=builder /app/helix-edge .
COPY --from=builder /app/helix-reactor .

EXPOSE 9000
EXPOSE 9100
EXPOSE 18080

CMD ["./helix-edge"]
