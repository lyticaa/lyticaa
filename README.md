# lyticaa-app

Lyticaa Dashboard (web/worker).

## Setup

### Golang

If you are new to Golang, please follow the setup instructions [here](https://golang.org/doc/install).

### Environment

Before running this project, please ensure that you have the following environment variables set:

```bash
APP_NAME=
AUTH0_CALLBACK_URI=
AUTH0_CLIENT_ID=
AUTH0_CLIENT_SECRET=
AUTH0_DOMAIN=
AUTH0_PASSWORD_RESET_URL=
AUTH0_URL=
AWS_ACCESS_KEY_ID=
AWS_REGION=
AWS_S3_UPLOAD_BUCKET=
AWS_SECRET_ACCESS_KEY=
AWS_SQS_QUEUE=
BASE_URL=
CLOUDAMQP_APIKEY=
CLOUDAMQP_QUEUE_EXPENSES_PENDING=
CLOUDAMQP_QUEUE_REPORTS_PUBLISHED=
CLOUDAMQP_URL=
DATABASE_URL=
ENV=
INTERCOM_ID=
MEMCACHED_PASSWORD=
MEMCACHED_SERVERS=
MEMCACHED_USERNAME=
NEW_RELIC_LICENSE_KEY=
PORT=3000
REDIS_PASSWORD=
REDIS_URL=
SENTRY_DSN=
SESSION_KEY=
STRIPE_ANNUAL_PLAN_ID=
STRIPE_ANNUAL_PRODUCT_ID=
STRIPE_CANCEL_URI=
STRIPE_MONTHLY_PLAN_ID=
STRIPE_MONTHLY_PRODUCT_ID=
STRIPE_PK=
STRIPE_SK=
STRIPE_SUCCESS_URI=
STRIPE_WHSEC=
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

### Run the Web App

The web app will require active Redis and Postgres instances. See below for how to start these. 

```bash
make run-web-service
```

The web app will then be accessible on http://localhost:3000.

### Run the Worker

```bash
make run-worker-service
```

The worker will then connect to RabbitMQ and listen for any incoming messages to process. You will need to ensure a local RabbitMQ instance is running (see below for how to start RabbitMQ).

## Sessions

Session data is stored in Redis. To start a local Redis instance, run:

```bash
make docker-redis
```

## Cache

Memcache is used to cache frequently requested objects. To start a local Memcache instance, run:

```bash
make docker-memcache
```

## Database

This project makes use of Postgres.

### Setup

To start a local Postgres instance, run:

```bash
make docker-pg
```

Then, to create the database and apply the correct role:

```bash
make create-database
make create-user
```

### Migrations

Add your migrations to the `db/migrations` folder. To apply the migrations:

```bash
make migrate
```

## Messaging

To communicate with other applications in the Lyticaa network, we use RabbitMQ.

To start a local RabbitMQ instance, run:

```bash
make docker-rabbitmq
```
