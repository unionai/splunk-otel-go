module github.com/unionai/splunk-otel-go/instrumentation/github.com/gomodule/redigo/splunkredigo

go 1.20

require (
	github.com/gomodule/redigo v1.8.9
	github.com/unionai/splunk-otel-go/instrumentation/internal v1.13.0
	github.com/stretchr/testify v1.8.4
	go.opentelemetry.io/otel v1.23.1
	go.opentelemetry.io/otel/trace v1.23.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.opentelemetry.io/otel/metric v1.23.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/unionai/splunk-otel-go/instrumentation/internal => ../../../../internal
