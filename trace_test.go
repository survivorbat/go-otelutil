package otelutil

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/trace"
)

//nolint:paralleltest // Can't be used because of the global state
func TestNewSpan_StartsSpanWithGivenName(t *testing.T) {
	// Arrange
	tracer := new(testTracer)

	Tracer = tracer
	t.Cleanup(func() { Tracer = nil })

	// Act
	_, _ = NewSpan(t.Context(), "testFunc", trace.WithSpanKind(20))

	// Assert
	assert.Equal(t, "testFunc", tracer.startCalledWithName)
	assert.Len(t, tracer.startCalledWithOpts, 1)
}

//nolint:paralleltest // Can't be used because of the global state
func TestNewSpan_StartsSpanWithReflectionIfNoNameGiven(t *testing.T) {
	// Arrange
	tracer := new(testTracer)

	Tracer = tracer
	t.Cleanup(func() { Tracer = nil })

	// Act
	_, _ = NewSpan(t.Context(), "", trace.WithSpanKind(15))

	// Assert
	assert.Equal(t, "go-otelutil.TestNewSpan_StartsSpanWithReflectionIfNoNameGiven", tracer.startCalledWithName)
	require.Len(t, tracer.startCalledWithOpts, 1)
}

func BenchmarkNewSpan_WithName(b *testing.B) {
	for b.Loop() {
		_, _ = NewSpan(b.Context(), "name")
	}
}

func BenchmarkNewSpan_WithoutName(b *testing.B) {
	for b.Loop() {
		_, _ = NewSpan(b.Context(), "")
	}
}

type testTracer struct {
	trace.Tracer

	startCalledWithName string
	startCalledWithOpts []trace.SpanStartOption
}

func (t *testTracer) Start(_ context.Context, name string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	t.startCalledWithName = name
	t.startCalledWithOpts = opts

	return nil, nil
}
