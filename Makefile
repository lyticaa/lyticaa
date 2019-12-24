GOBIN?=${GOPATH}/bin

all: lint install

lint-pre:
	@test -z $(gofmt -l .)
	@go mod verify

lint: lint-pre
	@golangci-lint run

lint-verbose: lint-pre
	@golangci-lint run -v

install: go.sum
	GO111MODULE=on go install -v ./cmd/dashd

clean:
	rm -f ${GOBIN}/{dashd}

tests:
	@go test -v -coverprofile .testCoverage.txt ./...

run-service:
	@dashd

run-stack:
	@docker-compose -f ./build/docker-compose.yml up --force-recreate --remove-orphans

pg:
	@docker-compose -f ./build/docker-compose.yml run --rm -p 5432:5432 --no-deps pg

create-user:
	PGPASSWORD=password psql -h localhost -U postgres -c "CREATE USER sellernomics WITH CREATEDB CREATEROLE PASSWORD 'password';"

create-database:
	 PGPASSWORD=password psql -h localhost -U postgres -c "CREATE DATABASE dashboard_development OWNER sellernomics;"

drop-database:
	PGPASSWORD=password psql -h localhost -U postgres -c "drop database dashboard_development;"

migrate:
	GO111MODULE=off CGO_ENABLED=1 CC=gcc go get -v github.com/rubenv/sql-migrate/...
	@sql-migrate ${DIRECTION}
