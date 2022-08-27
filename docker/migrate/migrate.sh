#!/bin/sh

/wait
/migrate \
  -source "github://kafugen/ocw-central-backend/migrations#main" \
  -database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}" \
  $@
