.PHONY: m.create m.up m.down m.version m.goto m.force

envfile := ./configs/.local.env

export $(cat ${envfile} | xargs -L1)

migrate_database := "sqlite://$(DB_CONN_STRING)?x-no-tx-wrap=true"

m.create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir ./migrations -seq -digits 4 $$name

m.up:
	@read -p "Enter N: " n; \
	migrate -path ./migrations -database $(migrate_database) up $$n

m.down:
	@read -p "Enter N: " n; \
	migrate -path ./migrations -database $(migrate_database) down $$n

m.version:
	@migrate -path ./migrations -database $(migrate_database) version 

m.goto:
	@read -p "Enter migration version: " version; \
	migrate -path ./migrations -database $(migrate_database) goto $$version

m.force:
	@read -p "Enter migration version: " version; \
	migrate -path ./migrations -database $(migrate_database) force $$version