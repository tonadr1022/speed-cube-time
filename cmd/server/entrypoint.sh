#!/bin/bash -e
# local if not specified

echo "[$(date)] Running DB migrations..."
migrate -database "${APP_DSN}" -path ./migrations up

echo "[$(date)] Starting server..."
server
