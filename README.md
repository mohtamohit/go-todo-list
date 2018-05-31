# go-todo-list
A simple todo list to learn about Postgres, Golang and development conventions for golang.

* [Dependency](#dependency)
* [Setup](#setup)
  + [Install go](#install-go)
* [Build](#build)
* [Start](#start)

## Dependency

- Postgresql 10
- GNU Make

## Setup

### Install go

- On OSX run brew install go.
- Follow instructions on https://golang.org/doc/install for other OSes.
- Make sure that the executable go is in your shell's path.
- Add the following in your .zshrc or .bashrc: (where <workspace_dir> is the directory in which you'll checkout your code).

``` bash
GOPATH=<workspace_dir>
export GOPATH
PATH="${PATH}:${GOPATH}/bin"
export PATH
```

## Build

1. Clone the repo to $GOPATH/src/github.com/mohtamohit/ directory.
1. Run `make setup` This installs prerequisites for the build.
1. Run `make build` This installs project dependencies and compiles the project to generate the binary.
1. Run `make db.setup`. This will run migrations.
1. Run `make test` to execute all tests.
1. You are now all set to start the app.
1. Also, if you're once into the app and feel like restarting the app with a new database, just run `make restart`.

## Start

1. Now run `out/go-todo web` to start the web app OR run `out/go-todo cli` to start the CLI app.
1. You can now visit `localhost:8080` to use the web app or if CLI then, simply continue with the terminal.