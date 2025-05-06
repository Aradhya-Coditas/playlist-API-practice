package tracer

import (
	"context"
	"omnenest-backend/src/utils/configs"

	"go.opentelemetry.io/otel"
	otlpTraceHTTP "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/otel/sdk/resource"
	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
)

var tracer trace.Tracer

// InitTracer initializes a tracer for the given service.
//
// Parameters:
// - ctx: context.Context
// - insecure: bool
// - serviceName: string
// - oltpEndpoint: string
// Return type(s):
// - *sdkTrace.TracerProvider
// - error
func InitTracer(ctx context.Context, insecure bool, serviceName string, oltpEndpoint string) (*sdkTrace.TracerProvider, error) {

	traceExporter, err := newOTLPExporter(ctx, oltpEndpoint, insecure)
	if err != nil {
		return nil, err
	}
	// trace provider
	tracerProvider := newTraceProvider(serviceName, traceExporter)

	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	tracer = tracerProvider.Tracer(serviceName)
	return tracerProvider, nil
}

// newOTLPExporter creates a new OTLP exporter.
//
// It takes the context, endpoint, and insecure flag as parameters and returns an sdkTrace.SpanExporter and an error.
func newOTLPExporter(ctx context.Context, endpoint string, insecure bool) (sdkTrace.SpanExporter, error) {
	options := []otlpTraceHTTP.Option{
		otlpTraceHTTP.WithEndpoint(endpoint),
	}

	if insecure {
		options = append(options, otlpTraceHTTP.WithInsecure())
	}

	return otlpTraceHTTP.New(ctx, options...)
}

// newTraceProvider creates a new TracerProvider.
//
// It takes a resource.Resource pointer and a sdkTrace.SpanProcessor as parameters.
// It returns a *sdkTrace.TracerProvider.
func newTraceProvider(serviceName string, exporter sdkTrace.SpanExporter) *sdkTrace.TracerProvider {
	tracerProvider := sdkTrace.NewTracerProvider(
		// sdkTrace.WithSampler(sdkTrace.AlwaysSample()),
		sdkTrace.WithBatcher(exporter),
		sdkTrace.WithSampler(sdkTrace.ParentBased(sdkTrace.TraceIDRatioBased(1.0))),
		sdkTrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
		)),
	)
	return tracerProvider
}

// GetTracer returns the initialized tracer.
func GetTracer() trace.Tracer {
	return tracer
}

func AddToSpan(spanCtx context.Context, methodName string) (context.Context, trace.Span) {
	applicationConfig := configs.GetApplicationConfig()
	isEnableTrace := applicationConfig.AppConfig.EnableOpenTelemetry
	if !isEnableTrace {
		return context.Background(), nil
	}
	childSpanCtx, span := GetTracer().Start(spanCtx, methodName)
	return childSpanCtx, span
}
