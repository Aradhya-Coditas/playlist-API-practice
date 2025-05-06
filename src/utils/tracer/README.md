# open telemetry package

This package provides functionality for initializing and managing OpenTelemetry tracers in a our backend application. It includes functions for setting up the tracer, creating spans, and retrieving the initialized tracer.

## Functions

### InitTracer

The function `InitTracer` Initializes a tracer for the given service using the OpenTelemetry protocol (OTLP).

```func InitTracer(ctx context.Context, insecure bool, serviceName string, oltpEndpoint string) (*sdkTrace.TracerProvider, error)```

- **Parameters**:
  - `ctx` : A context.Context object for the initialization.
  - `insecure` : A boolean indicating whether to use an insecure connection.
  - `serviceName` : A string representing the name of the service.
  - `oltpEndpoint` : A string specifying the endpoint for the OTLP exporter.


### newOTLPExporter

The function `newOTLPExporter` Creates a new OTLP exporter object.

```func newOTLPExporter(ctx context.Context, endpoint string, insecure bool) (sdkTrace.SpanExporter, error)```

- **Parameters**:
  - `ctx` : A context.Context object for the exporter creation.
  - `endpoint` : A string representing the endpoint for the OTLP exporter.
  - `insecure` : A boolean indicating whether to use an insecure connection.

### newTraceProvider

The function `newTraceProvider` Creates a new TracerProvider

```func newTraceProvider(serviceName string, exporter sdkTrace.SpanExporter) *sdkTrace.TracerProvider```

- **Parameters**:
  - `serviceName` : A string representing the name of the service.
  - `exporter` : A sdkTrace.SpanExporter object for exporting spans.

### GetTracer

The function `GetTracer` Retrieves the initialized tracer.

```func GetTracer() trace.Tracer```

### AddToSpan

The function `AddToSpan` Starts a new span, child span and adds it to the provided context.
```func AddToSpan(spanCtx context.Context, methodName string) (context.Context, trace.Span)```

Parameters:
spanCtx: A context.Context object representing the parent span context.
methodName: A string specifying the name of the method for which the span is created.


## Usage

To use this package, follow these steps:

1) Call `InitTracer` function to initialize the tracer with appropriate parameters.
2) Use `AddToSpan` function to add spans to your code where needed.
3) Optionally, use `GetTracer` function to retrieve the initialized tracer for any custom tracing needs.

Example usage:

```
import (
	"context"
	"your-package-path/tracer"
)

func main() {
	// Initialize the tracer
	_, err := tracer.InitTracer(context.Background(), false, "your-service-name", "http://your-otlp-endpoint:4318")
	if err != nil {
		// handle error
	}

	// Add spans to your code
	ctx := context.Background()
	ctx, span := tracer.AddToSpan(ctx, "your-method-name")
	defer span.End()

	// Your code here
}```