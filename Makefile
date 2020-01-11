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
	GO111MODULE=on go install -v ./cmd/workerd

clean:
	rm -f ${GOBIN}/{dashd}
	rm -f ${GOBIN}/{workerd}

tests:
	@go test -v -coverprofile .testCoverage.txt ./...

setup-yarn:
	yarn install

run-dashboard-service: build-assets
	@dashd

run-worker-service:
	@workerd

run-stack:
	@docker-compose -f ./build/docker-compose.yml up --force-recreate --remove-orphans

pg:
	@docker-compose -f ./build/docker-compose.yml run --rm -p 5432:5432 --no-deps pg

create-user:
	PGPASSWORD=password psql -h localhost -U postgres -c "CREATE USER lytica WITH CREATEDB CREATEROLE PASSWORD 'password';"

create-database:
	PGPASSWORD=password psql -h localhost -U postgres -c "CREATE DATABASE lytica_development OWNER lytica;"

drop-database:
	PGPASSWORD=password psql -h localhost -U postgres -c "drop database lytica_development;"

migrate:
	@go run tools/migrate/main.go

build-assets: setup-yarn

generate-docs: setup-yarn
	./node_modules/.bin/redoc-cli bundle ./api/docs/openapi.yml -o ./api/docs/index.html
