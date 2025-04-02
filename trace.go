package otelutil

import (
	"context"
	"path"
	"runtime"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// Tracer is used throughout this library to start new spans with. If it is unset at runtime, it will be set with
// an empty Tracer("") call.
var Tracer trace.Tracer

// NewSpan starts a new span using the Tracer with the given name. If the name of the span is empty,
// reflection will be used to determine the name of the calling function and use that as the span name.
//
// ⚠️  Not providing a name may be convenient, but comes at a performance penalty. Benchmarks in trace_test.go reveal that
// not providing the span name is ~5 times slower, which is relevant if milliseconds are of the essence.
//
// If the Tracer variable is nil, a new Tracer will be initialised with an empty name using otel.GetTracerProvider().
func NewSpan(ctx context.Context, name string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	if Tracer == nil {
		Tracer = otel.GetTracerProvider().Tracer("")
	}

	if name == "" {
		pc, _, _, _ := runtime.Caller(1)
		name = path.Base(runtime.FuncForPC(pc).Name())
	}

	//nolint:spancheck // It is returned, this is just a wrapper.
	return Tracer.Start(ctx, name, opts...)
}
