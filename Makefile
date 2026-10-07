include ./.env

DB_URL=postgres://$(DBUSER):$(DBPASS)@$(DBHOST):$(DBPORT)/$(DBNAME)?sslmode=disable
MIGRATION_PATH=db/migrations
SEEDER_PATH=db/seeds

migrate-create:
	@migrate create -ext sql -dir $(MIGRATION_PATH) -seq create_$(NAME)_table

migrate-up:
	@migrate -database $(DB_URL) -path $(MIGRATION_PATH) up

migrate-down:
	@migrate -database $(DB_URL) -path $(MIGRATION_PATH) down

print-db-url:
	@echo $(DB_URL)

seed-create:
	@powershell -Command "$$count = (Get-ChildItem -Path '$(SEEDER_PATH)' -ErrorAction SilentlyContinue).Count + 1; $$prefix = '{0:D3}' -f $$count; New-Item -Path '$(SEEDER_PATH)' -Name \"$${prefix}_$(NAME).sql\" -ItemType File"

# GNU MAKE - for seed create
seed-persons:
	@psql $(DB_URL) < $(SEEDER_PATH)/001_persons.sql