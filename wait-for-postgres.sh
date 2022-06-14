#!/bin/sh
# wait-for-postgres.sh

until psql -Atx $DB_CONNECTION_STRING -c '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing command"
exec "$@"