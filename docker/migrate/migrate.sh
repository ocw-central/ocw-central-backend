#!/bin/sh

/wait
/migrate \
  -source "github://shibukazu:${ACCESS_TOKEN}@kafugen/ocw-central-backend/migrations#main" \
  -database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}" up \
  $@

exit $?