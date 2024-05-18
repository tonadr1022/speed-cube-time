#!/bin/bash -e
# local if not specified

APP_ENV=${APP_ENV:-local}

CONFIG_FILE=./config/${APP_ENV}.yaml

ls -la
echo "up one dir"
ls -la ..

# if database string not specified, use config file
# sed finds and prints the string
if [[ -z ${APP_DSN} ]]; then
	export APP_DSN=$(sed -n 's/^dsn:[[:space:]]*"\(.*\)"/\1/p' "${CONFIG_FILE}")
fi

echo "[$(date)] Running DB migrations..."
cd ./migrations
cat 20240518173830_init.down.sql
cd ..

# migrate -path=./migrations -database "${APP_DSN}" version delete --version 20240518180957:
migrate -database "${APP_DSN}" -path ./migrations up

echo "[$(date)] Starting server..."

./server -config "${CONFIG_FILE}"
