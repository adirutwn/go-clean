#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE DATABASE myapp;
	GRANT ALL PRIVILEGES ON DATABASE myapp TO postgres;
EOSQL