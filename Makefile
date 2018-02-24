setup:
	curl https://glide.sh/get | sh
	glide install
	go get github.com/golang/mock/gomock
	go get github.com/golang/mock/mockgen

install:
	go install

db.create-migration:
	migrate create -ext sql -dir ./migration $(MIGRATION_NAME)