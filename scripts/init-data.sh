#!/bin/bash -e

PGPASSWORD=postgres psql -U postgres -h localhost -a -c "CREATE DATABASE bara;"
PGPASSWORD=postgres psql -U postgres -d bara -h localhost -a -f db/v1.sql
PGPASSWORD=postgres psql -U postgres -d bara -h localhost -a -f db/seeds_for_dev.sql

echo "Done"
