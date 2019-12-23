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
	@go test -mod=readonly -v -coverprofile .testCoverage.txt ./...

run-service:
	@dashd

run-stack:
	@docker-compose -f ./build/docker-compose.yml up --force-recreate --remove-orphans
