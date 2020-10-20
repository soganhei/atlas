#!/bin/sh

export TOKEN_SECRET="atlas" \
PORT=5000 \
DATABASE_PGSQL_URL="postgres://postgres:--PASSWORD--@localhost/--DB_NAME--?sslmode=disable"