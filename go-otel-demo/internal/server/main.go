package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-otel-demo/internal/config"
	"go-otel-demo/internal/tracing"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"io"
	"log"
	"net/http"
)

var globalCfg config.Config

func StartMain(cfg config.Config) {

	globalCfg = cfg

	stopTraceExporter := tracing.SetUp(cfg.Tracing)
	defer stopTraceExporter()

	router := mux.NewRouter()
	router.Use(
		otelmux.Middleware(cfg.Tracing.Service, otelmux.WithTracerProvider(otel.GetTracerProvider())),
	)

	router.HandleFunc("/vendors", getVendorsHandler)
	http.ListenAndServe(fmt.Sprintf(":%d", cfg.HttpPort), router)
}

func getVendorsHandler(w http.ResponseWriter, r *http.Request) {
	url := globalCfg.RailsAppURL

	log.Println("Request is forwarded to", url)

	ctx, span := otel.Tracer("").Start(r.Context(),
		"outbound.call", trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()

	// generate request
	request, _ := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/vendors", url), nil)

	client := &http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}

	res, err := client.Do(request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s", string(body))
}
