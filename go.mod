module gitlab.com/getlytica/lytica-app

replace gitlab.com/getlytica/lytica-app => ../getlytica/lytica-app

go 1.13

require (
	github.com/DATA-DOG/go-sqlmock v1.4.1
	github.com/aws/aws-sdk-go v1.31.7
	github.com/bufferapp/sqs-worker-go v0.0.0-20181101064454-7e780f286181
	github.com/coreos/go-oidc v2.2.1+incompatible
	github.com/garyburd/redigo v1.6.0 // indirect
	github.com/getsentry/sentry-go v0.6.1
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/golang-migrate/migrate/v4 v4.11.0
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.4
	github.com/gorilla/sessions v1.2.0
	github.com/heroku/x v0.0.24
	github.com/jmoiron/sqlx v1.2.0
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/lib/pq v1.5.2
	github.com/memcachier/mc v2.0.1+incompatible
	github.com/newrelic/go-agent v3.5.0+incompatible
	github.com/pquerna/cachecontrol v0.0.0-20180517163645-1555304b9b35 // indirect
	github.com/rs/zerolog v1.18.0
	github.com/stripe/stripe-go v70.15.0+incompatible
	github.com/stripe/stripe-go/v71 v71.15.0
	github.com/tealeg/xlsx v1.0.5
	github.com/urfave/negroni v1.0.0
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/text v0.3.2
	gopkg.in/boj/redistore.v1 v1.0.0-20160128113310-fc113767cd6b
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15
	gopkg.in/go-playground/validator.v9 v9.31.0
	gopkg.in/square/go-jose.v2 v2.5.1 // indirect
)
