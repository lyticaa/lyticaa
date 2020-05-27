module gitlab.com/getlytica/lytica-app

replace gitlab.com/getlytica/lytica-app => ../getlytica/lytica-app

go 1.13

require (
	github.com/DATA-DOG/go-sqlmock v1.4.0
	github.com/aws/aws-sdk-go v1.17.7
	github.com/bufferapp/sqs-worker-go v0.0.0-20181101064454-7e780f286181
	github.com/coreos/go-oidc v2.1.0+incompatible
	github.com/garyburd/redigo v1.6.0 // indirect
	github.com/getsentry/sentry-go v0.3.1
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/golang-migrate/migrate/v4 v4.7.1
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.3
	github.com/gorilla/sessions v1.2.0
	github.com/heroku/x v0.0.15
	github.com/jmoiron/sqlx v1.2.0
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/lib/pq v1.1.1
	github.com/mattn/go-sqlite3 v1.11.0 // indirect
	github.com/memcachier/mc v2.0.1+incompatible
	github.com/newrelic/go-agent v3.0.0+incompatible
	github.com/pquerna/cachecontrol v0.0.0-20180517163645-1555304b9b35 // indirect
	github.com/rs/zerolog v1.17.2
	github.com/stripe/stripe-go v68.4.0+incompatible
	github.com/tealeg/xlsx v1.0.5
	github.com/urfave/negroni v1.0.0
	golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550 // indirect
	golang.org/x/oauth2 v0.0.0-20191202225959-858c2ad4c8b6
	golang.org/x/sys v0.0.0-20191223224216-5a3cf8467b4e // indirect
	gopkg.in/boj/redistore.v1 v1.0.0-20160128113310-fc113767cd6b
	gopkg.in/go-playground/validator.v9 v9.31.0
	gopkg.in/square/go-jose.v2 v2.4.1 // indirect
	gopkg.in/yaml.v2 v2.2.7 // indirect
)
