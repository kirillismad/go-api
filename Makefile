migrate_database := "sqlite://$(DB_CONN_STRING)?x-no-tx-wrap=true"

.PHONY: m.create
m.create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir ./migrations -seq -digits 4 $$name

.PHONY: m.up
m.up:
	@read -p "Enter N: " n; \
	migrate -path ./migrations -database $(migrate_database) up $$n

.PHONY: m.down
m.down:
	@read -p "Enter N: " n; \
	migrate -path ./migrations -database $(migrate_database) down $$n

.PHONY: m.version
m.version:
	@migrate -path ./migrations -database $(migrate_database) version 

.PHONY: m.goto
m.goto:
	@read -p "Enter migration version: " version; \
	migrate -path ./migrations -database $(migrate_database) goto $$version

.PHONY: m.force
m.force:
	@read -p "Enter migration version: " version; \
	migrate -path ./migrations -database $(migrate_database) force $$version
