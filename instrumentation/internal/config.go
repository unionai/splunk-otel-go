// Copyright Splunk Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package internal

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
)

// Config contains configuration options.
type Config struct {
	// instName is the name of the instrumentation this Config is used for.
	instName string
	// Version is the version of the instrumentation this Config is used for.
	// It has to be set before user-provided options are applied.
	Version string

	Tracer     trace.Tracer
	Meter      metric.Meter
	Propagator propagation.TextMapPropagator

	DefaultStartOpts []trace.SpanStartOption
}

// NewConfig returns a Config for instrumentation with all options applied.
//
// If no TracerProvider or Propagator are specified with options, the default
// OpenTelemetry globals will be used.
func NewConfig(instrumentationName string, options ...Option) *Config {
	c := Config{instName: instrumentationName}

	for _, o := range options {
		if o != nil {
			o.Apply(&c)
		}
	}

	if c.Tracer == nil {
		c.Tracer = c.tracer(otel.GetTracerProvider())
	}
	if c.Meter == nil {
		c.Meter = c.meter(global.MeterProvider())
	}

	if c.Propagator == nil {
		c.Propagator = otel.GetTextMapPropagator()
	}

	return &c
}

// Copy returns a deep copy of c.
func (c *Config) Copy() *Config {
	newC := Config{
		instName:         c.instName,
		Tracer:           c.Tracer,
		Meter:            c.Meter,
		Propagator:       c.Propagator,
		DefaultStartOpts: make([]trace.SpanStartOption, len(c.DefaultStartOpts)),
	}

	copy(newC.DefaultStartOpts, c.DefaultStartOpts)

	return &newC
}

// ResolveTracer returns an OpenTelemetry tracer from the appropriate
// TracerProvider.
//
// If the passed context contains a span, the TracerProvider that created the
// tracer that created that span will be used. Otherwise, the TracerProvider
// from c is used.
func (c *Config) ResolveTracer(ctx context.Context) trace.Tracer {
	if span := trace.SpanFromContext(ctx); span.SpanContext().IsValid() {
		return c.tracer(span.TracerProvider())
	}
	return c.Tracer
}

// ResolveMeter returns an OpenTelemetry meter from the appropriate
// MeterProvider.
func (c *Config) ResolveMeter() metric.Meter {
	return c.Meter
}

// MergedSpanStartOptions returns a copy of opts with any DefaultStartOpts
// that c is configured with prepended.
func (c *Config) MergedSpanStartOptions(opts ...trace.SpanStartOption) []trace.SpanStartOption {
	if c == nil || len(c.DefaultStartOpts) == 0 {
		if len(opts) == 0 {
			return nil
		}
		cp := make([]trace.SpanStartOption, len(opts))
		copy(cp, opts)
		return cp
	}

	merged := make([]trace.SpanStartOption, len(c.DefaultStartOpts)+len(opts))
	copy(merged, c.DefaultStartOpts)
	copy(merged[len(c.DefaultStartOpts):], opts)
	return merged
}

// WithSpan wraps the function f with a span named name.
func (c *Config) WithSpan(ctx context.Context, name string, f func(context.Context) error, opts ...trace.SpanStartOption) error {
	sso := c.MergedSpanStartOptions(opts...)
	ctx, span := c.ResolveTracer(ctx).Start(ctx, name, sso...)
	err := f(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}
	span.End()

	return err
}

// tracer creates a tracer using the passed TracerProvider.
func (c *Config) tracer(tp trace.TracerProvider) trace.Tracer {
	opts := []trace.TracerOption{trace.WithSchemaURL(semconv.SchemaURL)}
	if c.Version != "" {
		opts = append(opts, trace.WithInstrumentationVersion(c.Version))
	}
	return tp.Tracer(c.instName, opts...)
}

// meter creates a meter using the passed MeterProvider.
func (c *Config) meter(mp metric.MeterProvider) metric.Meter {
	opts := []metric.MeterOption{metric.WithSchemaURL(semconv.SchemaURL)}
	if c.Version != "" {
		opts = append(opts, metric.WithInstrumentationVersion(c.Version))
	}
	return mp.Meter(c.instName, opts...)
}
