FROM golang:1.20-alpine

WORKDIR /app

COPY . c
WORKDIR /app/c

COPY go.mod .
COPY go.sum .
RUN go mod download


RUN CGO_ENABLED=0 go build -o /mikrotik-exporter


FROM scratch
WORKDIR /

EXPOSE 9436

COPY scripts/start.sh /
COPY --from=0 /mikrotik-exporter /

ENTRYPOINT ["/start.sh"]