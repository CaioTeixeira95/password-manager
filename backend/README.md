# Running

Be sure that you've installed [docker](https://docs.docker.com/desktop/) and [docker-compose](https://docs.docker.com/compose/install/) or/and [golang 1.20](https://go.dev/doc/install).

## Using docker compose

```sh
$ docker-compose up --build
```

## Using golang

```sh
$ go run main.go
```

# Tests

```sh
$ go test -v -race -cover ./...
```

# Algorith & Approach

## Libraries

- [Fiber](https://docs.gofiber.io/): A lightweight web framework that provides some useful tools to handle HTTP requests.
- [Testify](https://github.com/stretchr/testify): A awesome testing library.

## Architecture

The project architecture is divided in `model`, `repository`, `service`, and `serve`.

- [model](./model/): The models represents the project's model.
- [repository](./repository/): This layer has the responsibility of communicating with the storage service - in this case we store in the memory.
- [service](./service/): Here is where the business rules lives and can be reused independent of the context.
- [serve](./serve/): The transport layer and where the HTTP handlers live.

Each layer requires its own dependencies this way it's easy to test and change components.
