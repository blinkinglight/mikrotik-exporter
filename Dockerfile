FROM debian:9.9-slim

EXPOSE 9436

COPY scripts/start.sh /app/
COPY dist/github.com/blinkinglight/mikrotik-exporter_linux_amd64 /app/github.com/blinkinglight/mikrotik-exporter

RUN chmod 755 /app/*

ENTRYPOINT ["/app/start.sh"]