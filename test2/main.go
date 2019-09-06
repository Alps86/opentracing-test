package main

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmhttp"
	"go.elastic.co/apm/module/apmot"
)

func main() {
	opentracing.SetGlobalTracer(apmot.New())

	defer apm.DefaultTracer.Flush(nil)

	tracingId := "00-66562ed6afda0a9b48e7a14852800279-66562ed6afda0a9b-01"

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

	writer := opentracing.TextMapCarrier{}
	err = opentracing.GlobalTracer().Inject(
		serverSpan.Context(),
		opentracing.TextMap,
		writer,
	)

	if err != nil {
		panic(err)
	}

	tracingId = writer[apmhttp.TraceparentHeader]
	fmt.Println(tracingId)

	defer serverSpan.Finish()
}
