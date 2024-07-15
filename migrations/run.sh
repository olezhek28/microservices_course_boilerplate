#!/bin/bash

export MIGRATION_DIR=./migrations
export MIGRATION_DSN="host=$PG_HOST port=$PG_PORT dbname=$PG_DBNAME user=$PG_USER password=$PG_PWD sslmode=disable"

sleep 2 && goose -dir "${MIGRATION_DIR}" postgres "${MIGRATION_DSN}" up -v