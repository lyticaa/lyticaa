module github.com/lyticaa/lyticaa

replace github.com/lyticaa/lyticaa => ../lyticaa/lyticaa

go 1.15

require (
	github.com/aws/aws-sdk-go v1.40.46
	github.com/bufferapp/sqs-worker-go v0.0.0-20181101064454-7e780f286181
	github.com/coreos/go-oidc v2.2.1+incompatible
	github.com/garyburd/redigo v1.6.0 // indirect
	github.com/getsentry/sentry-go v0.7.0
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/golang-migrate/migrate/v4 v4.11.0
	github.com/gorilla/csrf v1.7.0
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/sessions v1.2.1
	github.com/heroku/x v0.0.26
	github.com/jmoiron/sqlx v1.2.0
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/lib/pq v1.8.0
	github.com/memcachier/mc v2.0.1+incompatible
	github.com/newrelic/go-agent v3.9.0+incompatible
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/pquerna/cachecontrol v0.0.0-20180517163645-1555304b9b35 // indirect
	github.com/rs/zerolog v1.20.0
	github.com/stretchr/testify v1.6.1
	github.com/stripe/stripe-go/v72 v72.23.0
	github.com/tealeg/xlsx v1.0.5
	github.com/urfave/negroni v1.0.0
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/text v0.3.6
	gopkg.in/boj/redistore.v1 v1.0.0-20160128113310-fc113767cd6b
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f
	gopkg.in/go-playground/validator.v9 v9.31.0
	gopkg.in/square/go-jose.v2 v2.5.1 // indirect
	syreclabs.com/go/faker v1.2.2
)
