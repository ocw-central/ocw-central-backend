#!/bin/sh

/wait

if [ $APP_ENV = "LOCAL" ]; then
  /migrate \
    -path $MIGRATIONS_DIR \
    -database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}" up
elif [ $APP_ENV = "PROD" -o $APP_ENV = "DEV" ]; then
  /migrate \
    -source "github://${ACCESS_USER}:${ACCESS_TOKEN}@kafugen/ocw-central-backend/migrations#${MIGRATE_BRANCH}" \
    -database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}" up
else
  echo "[INFO migrate] Unknown APP_ENV: $APP_ENV"
  exit 1
fi

if [ $? -eq 0 ]; then
  echo "[INFO migrate] Migrations succeeded."
else
  echo "[INFO migrate] Migrations failed."
  exit 1
fi

./main
