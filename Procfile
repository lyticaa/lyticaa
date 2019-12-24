release: GO111MODULE=off go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate && ./bin/migrate -source file://db/migrations/ -database ${DATABASE_URL} up
web: bin/dashd
