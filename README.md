# dashboard

Seller Nomics Dashboard

## Setup

### Golang

If you are new to Golang, please follow the setup instructions [here](https://golang.org/doc/install).

## Setup

### Environment

Before running this project, please ensure that you have the following environment variables set:

```bash
APP_NAME=
SENTRY_DSN=
NEWRELIC_LICENSE_KEY=
PORT=
SESSION_KEY=
REDIS_URL=
REDIS_PASSWORD=
AUTH0_URL=
AUTH0_CLIENT_ID=
AUTH0_CLIENT_SECRET=
AUTH0_CALLBACK_URL=
DB_HOST=
DB_PORT=
DB_USERNAME=
DB_PASSWORD=
DB_NAME=
DB_SSLMODE=
```

If you are unsure as to what these values ought to be, then please check with a colleague.

### Linter

To run the linter:

```bash
make lint
```

### Tests

To run the tests and see test coverage:

```bash
make tests
```

### Install

To compile and install the binary:

```bash
make install
```

### Run the Service

```bash
make run-service
```

## Database

This project makes use of Postgres.

### Setup

To start a local (Docker) instance of Postgres, simply run:

```bash
make pg
```

Then, to create the database and apply the correct role:

```bash
make create-database
make create-user
```

### Migrations

Install the [migrate](https://github.com/golang-migrate/migrate) tool:

```bash
make sql-migrate
```

Add your migrations to the `db/migrations` folder and then run:

```bash
make DIRECTION=up migrate-database
```

To rollback your recent migrations, run:

```bash
make DIRECTION=down migrate-database
```

## Docker

A Docker stack is provided with this project. To boot the stack, simply run:

```bash
make run-stack
```

Please ensure that prior to running this, you add the above environment variables to the `build/.env` file. Docker Compose will use these when building the container.
