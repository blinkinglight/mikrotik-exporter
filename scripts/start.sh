#!/bin/sh

if [ ! -x /app/github.com/blinkinglight/mikrotik-exporter ]; then
  chmod 755 /app/github.com/blinkinglight/mikrotik-exporter
fi

if [ -z "$CONFIG_FILE" ]
then
    /app/github.com/blinkinglight/mikrotik-exporter -device $DEVICE -address $ADDRESS -user $USER -password $PASSWORD
else
    /app/github.com/blinkinglight/mikrotik-exporter -config-file $CONFIG_FILE
fi
