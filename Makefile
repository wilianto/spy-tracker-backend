setup:
	curl https://glide.sh/get | sh
	glide install

install:
	go install

db.create-migration:
	migrate create -ext sql -dir ./migration $(MIGRATION_NAME)