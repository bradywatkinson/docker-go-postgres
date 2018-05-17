#!/bin/bash
set -e

echo "Initialising: postgresql://test:test@postgres/test"

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
    CREATE USER test;
    CREATE DATABASE test;
    GRANT ALL PRIVILEGES ON DATABASE test TO test;
EOSQL
