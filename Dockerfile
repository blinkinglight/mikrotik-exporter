FROM golang:1.20-alpine

WORKDIR /app

COPY . c
WORKDIR /app/c

COPY go.mod .
COPY go.sum .
RUN go mod download


RUN CGO_ENABLED=0 go build -o /mikrotik-exporter


FROM alpine
WORKDIR /

EXPOSE 9436

COPY --from=0 /mikrotik-exporter /
COPY scripts/start.sh /start.sh

ENTRYPOINT ["/start.sh"]