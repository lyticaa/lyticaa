module gitlab.com/sellernomics/dashboard

replace gitlab.com/sellernomics/dashboard => ../sellernomics/dashboard

go 1.13

require (
	github.com/coreos/go-oidc v2.1.0+incompatible
	github.com/garyburd/redigo v1.6.0 // indirect
	github.com/getsentry/sentry-go v0.3.1
	github.com/gorilla/mux v1.7.3
	github.com/gorilla/sessions v1.2.0
	github.com/newrelic/go-agent v3.0.0+incompatible
	github.com/pquerna/cachecontrol v0.0.0-20180517163645-1555304b9b35 // indirect
	github.com/rs/zerolog v1.17.2
	github.com/thedevsaddam/renderer v1.2.0
	github.com/urfave/negroni v1.0.0
	golang.org/x/oauth2 v0.0.0-20191202225959-858c2ad4c8b6
	gopkg.in/boj/redistore.v1 v1.0.0-20160128113310-fc113767cd6b
	gopkg.in/square/go-jose.v2 v2.4.1 // indirect
)
