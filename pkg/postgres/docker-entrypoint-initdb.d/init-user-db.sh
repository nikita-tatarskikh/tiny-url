#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB"  <<-EOSQL
    CREATE USER "tiny-user" WITH PASSWORD 'tiny-password';
    CREATE DATABASE "tiny";
    GRANT ALL PRIVILEGES ON DATABASE "tiny" TO "tiny-user";
EOSQL