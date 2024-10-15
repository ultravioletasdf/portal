#!/bin/bash

# Load the .env file
if [[ -f .env ]]; then
    export $(grep -v '^#' .env | xargs)
else
    echo ".env file not found."
    exit 1
fi
# Check if variables are set
if [[ -z "$TURSO_ADDRESS" || -z "$TURSO_TOKEN" ]]; then
    echo "TURSO_ADDRESS or TURSO_TOKEN is not set in the .env file."
    exit 1
fi
TURSO_ADDRESS="${TURSO_ADDRESS/libsql:\/\//libsql+wss:\/\/}"
TURSO_URL="${TURSO_ADDRESS}?authToken=${TURSO_TOKEN}"
atlas schema apply --url $TURSO_URL --dev-url "sqlite://dev?mode=memory" --to "file://sql/schema.sql" --exclude "_litestream" 
