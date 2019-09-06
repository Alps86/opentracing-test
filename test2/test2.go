package main

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmhttp"
	"go.elastic.co/apm/module/apmot"
)

func main() {
	opentracing.SetGlobalTracer(apmot.New())

	defer apm.DefaultTracer.Flush(nil)

	tracingId := "00-d305d61d2084153438420e8e5ef2aad1-d305d61d20841534-01"

	carrier := opentracing.TextMapCarrier{}
	carrier[apmhttp.TraceparentHeader] = tracingId

	wireContext, err := opentracing.GlobalTracer().Extract(
		opentracing.TextMap,
		carrier,
	)
	if err != nil {
		// Optionally record something about err here
	}

	serverSpan := opentracing.StartSpan(
		"test",
		ext.RPCServerOption(wireContext))

	defer serverSpan.Finish()
}
