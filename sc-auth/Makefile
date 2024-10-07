db-schema: ## Generate db schema
	sqlc generate

db-schema-up: ## Migrate up
	migrate -path sc-auth/internal/db/migration -database 'postgresql://root:secret@localhost:5432/sc_db?sslmode=disable' up

db-schema-down: ## Migrate down
	migrate -path sc-auth/internal/db/migration -database 'postgresql://root:secret@localhost:5432/sc_db?sslmode=disable' down