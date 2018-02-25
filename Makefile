setup:
	curl https://glide.sh/get | sh
	glide install

install:
	go install

db.migrate:
	migrate -source file://migration -database $(DB_USER)://$(DB_HOST):$(DB_PORT)/$(DB_NAME)_$(ENV)?sslmode=disable up $(NUMBER)

db.rollback:
	migrate -source file://migration -database $(DB_USER)://$(DB_HOST):$(DB_PORT)/$(DB_NAME)_$(ENV)?sslmode=disable down $(NUMBER)

db.create:
	createdb -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -W $(DB_NAME)_$(ENV)

db.create-migration:
	migrate create -ext sql -dir ./migration $(MIGRATION_NAME)