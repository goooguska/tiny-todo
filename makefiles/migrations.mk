DB_MIGRATIONS_SOURCE=./db/migrations
DATABASE_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)
DB_SCHEMA = public
MIGRATE_APP_PATH = migrate

create-migration:
	$(MIGRATE_APP_PATH) create -ext sql -dir $(DB_MIGRATIONS_SOURCE)/$(DB_SCHEMA) -seq $(NAME)

migrate:
	$(MIGRATE_APP_PATH) -database $(DATABASE_URL) -path $(DB_MIGRATIONS_SOURCE)/$(DB_SCHEMA) up

rollback:
	$(MIGRATE_APP_PATH) -database $(DATABASE_URL) -path $(DB_MIGRATIONS_SOURCE)/$(DB_SCHEMA) down