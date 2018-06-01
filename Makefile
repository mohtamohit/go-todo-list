.PHONY : all
all: build-deps build test

ALL_PACKAGES = $(shell go list ./...)
UNIT_TEST_PACKAGES = $(shell go list ./...)
DB_NAME = "todo_dev"
TEST_DB_NAME = "todo_test"
TEST_DB_PORT = 5432
DB_PORT = 5432
APP_EXECUTEABLE = "out/todo"

setup:
	go get -u github.com/golang/dep/cmd/dep

build-deps:
	dep ensure

update-deps:
	dep ensure

compile: 
	mkdir -p out/
	go build -o $(APP_EXECUTEABLE) 

build: build-deps compile fmt vet

install:
	go install ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

db.setup: db.create db.migrate

db.create:
	createdb -p $(DB_PORT) -Opostgres -Eutf8 $(DB_NAME)

db.migrate:
	$(APP_EXECUTEABLE) migrate

db.drop:
	dropdb -p $(DB_PORT) --if-exists -Upostgres $(DB_NAME)

db.reset: db.drop db.create db.migrate

db.rollback:
	$(APP_EXECUTABLE) rollback

test: testdb.reset
	ENVIRONMENT=test go test $(UNIT_TEST_PACKAGES) 

testdb.reset: testdb.drop testdb.create testdb.migrate

testdb.drop:
	dropdb -p $(TEST_DB_PORT) --if-exists -Upostgres $(TEST_DB_NAME)

testdb.create:
	createdb -p $(TEST_DB_PORT) -Opostgres -Eutf8 $(TEST_DB_NAME)

testdb.migrate:
	ENVIRONMENT=test $(APP_EXECUTEABLE) migrate

testdb.rollback:
	ENVIRONMENT=test $(APP_EXECUTEABLE) rollback

restart: build db.drop db.setup