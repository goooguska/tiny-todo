-include .env

DOCKER-COMPOSE-COMMAND=docker compose

start-db:
	$(DOCKER-COMPOSE-COMMAND) --env-file .env -f docker/docker-compose.db.yaml up -d

stop-db:
	$(DOCKER-COMPOSE-COMMAND) --env-file .env -f docker/docker-compose.db.yaml stop

include makefiles/migrations.mk