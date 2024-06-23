#!/bin/bash -e
# local if not specified

echo "[$(date)] Running DB migrations..."
migrate -database "${DATABASE_URL}" -path ./migrations up

echo "[$(date)] Starting server..."
server
