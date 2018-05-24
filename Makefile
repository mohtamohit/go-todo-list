.PHONY : all
all: build-deps build test

ALL_PACKAGES = $(shell go list ./...)

DB_NAME = "todo_db"
TEST_DB_NAME = "todo_db_test"
TEST_DB_PORT = 5432
DB_PORT = 5432

setup:
	go get -u github.com/golang/dep/cmd/dep

build-deps:
	dep ensure

update-deps:
	dep ensure

build: build-deps compile fmt vet

install:
	go install ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

