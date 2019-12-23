module gitlab.com/sellernomics/dashboard

replace gitlab.com/sellernomics/dashboard => ../sellernomics/dashboard

go 1.13

require (
	github.com/getsentry/sentry-go v0.3.1
	github.com/gorilla/mux v1.7.3
	github.com/newrelic/go-agent v3.0.0+incompatible
	github.com/rs/zerolog v1.17.2
	github.com/thedevsaddam/renderer v1.2.0
)
