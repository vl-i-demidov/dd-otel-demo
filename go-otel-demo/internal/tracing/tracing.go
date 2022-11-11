package tracing

import (
	"context"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"log"
	"time"
)

func SetUp(cfg Config) (stop func()) {

	client := otlptracegrpc.NewClient(
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(cfg.ReceiverEndpoint))

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	traceExp, err := otlptrace.New(ctx, client)
	if err != nil {
		log.Fatal(err)
	}

	res, err := resource.New(ctx,
		resource.WithAttributes(
			attribute.String(ext.Environment, cfg.Environment),
			attribute.String(ext.ServiceName, cfg.Service),
			attribute.String(ext.Version, cfg.AppVersion),
		))

	tracerProvider := tracesdk.NewTracerProvider(
		tracesdk.WithSampler(tracesdk.AlwaysSample()),
		tracesdk.WithBatcher(traceExp),
		tracesdk.WithResource(res),
	)

	compositePropagator := propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, b3.New(b3.WithInjectEncoding(b3.B3MultipleHeader)))
	otel.SetTextMapPropagator(compositePropagator)
	otel.SetTracerProvider(tracerProvider)

	return func() {
		cxt, cancel := context.WithTimeout(ctx, 50*time.Second)
		defer cancel()

		if err := traceExp.Shutdown(cxt); err != nil {
			otel.Handle(err)
		}
	}
}
