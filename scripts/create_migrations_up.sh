#!/bin/bash

go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

SERVICES=("sc-auth" "sc-user" "sc-post" "sc-notification")

for SERVICE_NAME in "${SERVICES[@]}"; do
  MIGRATION_PATH="$SERVICE_NAME/internal/db/migration"

  echo "$SERVICE_NAME"
  echo "migration path: $MIGRATION_PATH"
  migrate -database "postgresql://root:secret@localhost:5432/sc_db?sslmode=disable&x-migrations-table=migration_$SERVICE_NAME" \
          -path "$MIGRATION_PATH" \
          -verbose up
done