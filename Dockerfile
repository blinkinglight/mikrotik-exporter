FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 go build -o /mikrotik-exporter


FROM scratch
WORKDIR /

EXPOSE 9436

COPY scripts/start.sh /app/
COPY --from=0 /mikrotik-exporter /app

ENTRYPOINT ["/app/start.sh"]