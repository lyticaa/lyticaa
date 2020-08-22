GOBIN?=${GOPATH}/bin

all: lint install

lint-pre:
	@test -z $(gofmt -l .)
	@go mod verify

lint: lint-pre
	@golangci-lint run

lint-verbose: lint-pre
	@golangci-lint run -v --timeout=5m

install: go.sum
	GO111MODULE=on go install -v ./cmd/webd
	GO111MODULE=on go install -v ./cmd/workerd

clean:
	rm -f ${GOBIN}/{webd,workerd}

tests:
	@go test -v -coverprofile .testCoverage.txt ./...

setup-yarn:
	yarn install

build-assets: setup-yarn

run-web-service: build-assets
	@webd

run-worker-service:
	@workerd

docker-pg:
	@docker-compose -f ./build/docker-compose.yml run --rm -p 5432:5432 --no-deps pg

docker-redis:
	@docker-compose -f ./build/docker-compose.yml run --rm -p 6379:6379 --no-deps -d redis

docker-memcache:
	@docker-compose -f ./build/docker-compose.yml run --rm -p 11211:11211 --no-deps -d memcache

create-user:
	PGPASSWORD=password psql -h localhost -U postgres -c "CREATE USER lyticaa WITH SUPERUSER PASSWORD 'password';"

create-database:
	PGPASSWORD=password psql -h localhost -U postgres -c "CREATE DATABASE lyticaa_development OWNER lyticaa;"

drop-database:
	PGPASSWORD=password psql -h localhost -U postgres -c "drop database lyticaa_development;"

migrate:
	@go run cmd/migrate/main.go

generate-docs: setup-yarn
	./node_modules/.bin/redoc-cli bundle ./api/docs/openapi.yml -o ./api/docs/index.html
