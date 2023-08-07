# ArtVortex API

This is the API for ArtVortex project. It's built with Go and uses GraphQL to interact with the client-side. It also has Postgres as a database backend.

## Installation

### Prerequisites

Make sure the following software is installed on your system:

- PGAdmin
- Postgres: 15.2
- [golangci-lint](https://golangci-lint.run/usage/install/#local-installation)
- [Geth](https://geth.ethereum.org/docs/getting-started/installing-geth)
- Kubo CLI - IPFS
- docker

### Instructions

1. Clone the repository and navigate to the project root directory.

2. Install the project dependencies using the following command:

```sh
make dep
```

3. Buid (compile server into a binary file )

```sh
make build
```

4. There are some default service configs setup in system, if need you can override them with yours (especially the db's configuration)

```
export DB_HOST=locahost
export DB_PORT=5432
export DB_USER=pguser
export DB_NAME=artvortex
export DB_PASSWORD=
export HOST=http://locahost
export PORT=8000
```

5. Start server (assume db is already running)

```
make run
```

or to build and run server:

```
make start
```

6. Clean binary

```
make clean
```

**Using docker**
```
docker compose up
```

## Packages

This project uses the following packages:

- [GORM] (https://gorm.io/)
- [echo] (https://echo.labstack.com/)
- [namsral flag] (github.com/namsral/flag)
- [gqlgen](https://github.com/99designs/gqlgen)
- [GoDotEnv] (https://github.com/joho/godotenv)
- [Go postgres driver](https://github.com/go-gorm/postgres)
- [go-ipfs-api](https://github.com/ipfs/go-ipfs-api)

## Example Query and Mutation

Open http://localhost:8000/playground

```sh
mutation createUser {
  createUser(
    input: {UserName: "user01_test", Email: "user01_test@gmail.com"}
  ) {
    UserName
    Email
    CreditBalance
    CreatedAt
  }
}

query users {
  users{
    UserName
    Email
  }
}
```

then a user `user01_test` with email `user01_test@gmail.com` will be added into your Postgres Users table
