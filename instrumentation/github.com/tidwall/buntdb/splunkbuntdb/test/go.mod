module github.com/unionai/splunk-otel-go/instrumentation/github.com/tidwall/buntdb/splunkbuntdb/test

go 1.20

require (
	github.com/unionai/splunk-otel-go/instrumentation/github.com/tidwall/buntdb/splunkbuntdb v1.13.0
	github.com/stretchr/testify v1.8.4
	github.com/tidwall/buntdb v1.3.0
	go.opentelemetry.io/otel v1.23.1
	go.opentelemetry.io/otel/sdk v1.23.1
	go.opentelemetry.io/otel/trace v1.23.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/unionai/splunk-otel-go/instrumentation/internal v1.13.0 // indirect
	github.com/tidwall/btree v1.7.0 // indirect
	github.com/tidwall/gjson v1.17.1 // indirect
	github.com/tidwall/grect v0.1.4 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/tidwall/rtred v0.1.2 // indirect
	github.com/tidwall/tinyqueue v0.1.1 // indirect
	go.opentelemetry.io/otel/metric v1.23.1 // indirect
	golang.org/x/sys v0.17.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/unionai/splunk-otel-go/instrumentation/github.com/tidwall/buntdb/splunkbuntdb => ../
	github.com/unionai/splunk-otel-go/instrumentation/internal => ../../../../../internal
)
