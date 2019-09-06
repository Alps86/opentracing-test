package main

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmhttp"
	"go.elastic.co/apm/module/apmot"
)

func main() {
	opentracing.SetGlobalTracer(apmot.New())
	defer apm.DefaultTracer.Flush(nil)

	span := opentracing.StartSpan("parent")

	writer := opentracing.TextMapCarrier{}
	err := opentracing.GlobalTracer().Inject(
		span.Context(),
		opentracing.TextMap,
		writer,
	)

	if err != nil {
		panic(err)
	}

	tracingId := writer[apmhttp.TraceparentHeader]
	fmt.Println(tracingId)

	span.Finish()
}
