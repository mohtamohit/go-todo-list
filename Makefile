.PHONY : all
all: build-deps build test

ALL_PACKAGES = $(shell go list ./...)
UNIT_TEST_PACKAGES = $(shell go list ./...)
DB_NAME = "todo_db"
TEST_DB_NAME = "todo_test_db"
TEST_DB_PORT = 5432
DB_PORT = 5432
APP_EXECUTEABLE = "out/go-todo"

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

test: testdb.reset
	ENVIRONMENT=test go test $(UNIT_TEST_PACKAGES) 

testdb.reset: testdb.drop testdb.create testdb.migrate

testdb.drop:
	dropdb -p $(TEST_DB_PORT) --if-exists -Upostgres $(TEST_DB_NAME)

testdb.create:
	createdb -p $(TEST_DB_PORT) -Opostgres -Eutf8 $(TEST_DB_NAME)

testdb.migrate:
	ENVIRONMENT=TEST $(APP_EXECUTEABLE) migrate


