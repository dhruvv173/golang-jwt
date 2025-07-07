#!/bin/sh
set -e

host="$DB_HOST"
port="$DB_PORT"

until pg_isready -h "$host" -p "$port" > /dev/null 2>&1; do
  echo "Waiting for Postgres at $host:$port..."
  sleep 2
done

echo "Postgres is up, starting the app..."
exec "$@"
