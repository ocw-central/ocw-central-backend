#!/bin/sh

/wait
/migrate \
  -path $MIGRATIONS_DIR \
  -database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}" \
  $@
