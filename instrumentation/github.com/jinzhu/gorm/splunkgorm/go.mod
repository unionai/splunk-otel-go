module github.com/signalfx/splunk-otel-go/instrumentation/github.com/jinzhu/gorm/splunkgorm

go 1.18

require (
	github.com/jinzhu/gorm v1.9.16
	github.com/signalfx/splunk-otel-go/instrumentation/database/sql/splunksql v1.4.0
	github.com/stretchr/testify v1.8.2
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/signalfx/splunk-otel-go v1.4.0 // indirect
	github.com/signalfx/splunk-otel-go/instrumentation/internal v1.4.0 // indirect
	go.opentelemetry.io/otel v1.14.0 // indirect
	go.opentelemetry.io/otel/metric v0.37.0 // indirect
	go.opentelemetry.io/otel/trace v1.14.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/signalfx/splunk-otel-go => ../../../../../
	github.com/signalfx/splunk-otel-go/instrumentation/database/sql/splunksql => ../../../../database/sql/splunksql
	github.com/signalfx/splunk-otel-go/instrumentation/internal => ../../../../internal/
)
