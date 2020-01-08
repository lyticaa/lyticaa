# lytica

Main Lytica app (dashboard/worker).

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
STRIPE_PK=
STRIPE_SK=
AWS_REGION=
AWS_ACCESS_KEY_ID=
AWS_SECRET_ACCESS_KEY=
AWS_SQS_QUEUE=
AWS_S3_UPLOAD_BUCKET=
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

To compile and install the binaries:

```bash
make install
```

### Run the Dashboard

```bash
make run-dashboard-service
```

The dashboard will then be accessible on http://localhost:3000.

### Run the Worker

```bash
make run-worker-service
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

Add your migrations to the `db/migrations` folder. To apply the migrations, simply run:

```bash
make migrate
```

## Docker

A Docker stack is provided with this project. To boot the stack, simply run:

```bash
make run-stack
```

Please ensure that prior to running this, you add the above environment variables to the `build/.env` file. Docker Compose will use these when building the container.
