FROM golang:alpine AS builder

WORKDIR /app
COPY . /app

RUN go build -o ntfy2gotify ntfy2gotify.go


FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/ntfy2gotify .
COPY --from=builder /app/config.json .

RUN chmod +x ntfy2gotify
CMD ["./ntfy2gotify"]