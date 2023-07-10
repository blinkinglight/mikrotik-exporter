#!/bin/sh

if [ ! -x /app/mikrotik-exporter ]; then
  chmod 755 /mikrotik-exporter
fi

if [ -z "$CONFIG_FILE" ]
then
    /mikrotik-exporter -device $DEVICE -address $ADDRESS -user $USER -password $PASSWORD
else
    /mikrotik-exporter -config-file $CONFIG_FILE
fi
