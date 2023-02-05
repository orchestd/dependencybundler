package trace

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/orchestd/transport/client"
	"net/http"
	"net/http/httputil"
	"reflect"
	"unsafe"
)

// TracerRESTClientInterceptor is a REST tracing client interceptor, it can log req/resp if needed
func TracerRESTClientInterceptor(deps tracingDeps) client.HTTPClientInterceptor {
	return func(req *http.Request, handler client.HTTPHandler) (resp *http.Response, err error) {
		if deps.Tracer == nil {
			return handler(req)
		}
		span, ctx := deps.newClientSpanForREST(req)
		defer span.Finish()

		req = req.WithContext(ctx)
		resp, err = handler(req)
		if err != nil {
			ext.LogError(span, err)
		} else {
			ext.HTTPStatusCode.Set(span, uint16(resp.StatusCode))
			if respDump, dumpErr := httputil.DumpResponse(resp, true); dumpErr == nil {
				addBodyToSpan(span, "response", respDump)
			} else {
				deps.Logger.WithError(dumpErr).Debug(ctx, "failed to dump response")
			}
		}

		return
	}
}

func (d tracingDeps) newClientSpanForREST(req *http.Request) (opentracing.Span, context.Context) {
	var ctx = context.Background()
	if req.Context() != nil {
		ctx = req.Context()
	}

	span, clientContext := opentracing.StartSpanFromContextWithTracer(ctx, d.Tracer, req.URL.Path, ext.SpanKindRPCClient, httpTag)
	if reqDump, dumpErr := httputil.DumpRequestOut(req, true); dumpErr == nil {
		addBodyToSpan(span, "request", reqDump)
	} else {
		d.Logger.WithError(dumpErr).Debug(ctx, "failed to dump request")
	}

	ext.HTTPUrl.Set(span, fmt.Sprintf("%v", req.URL))
	ext.HTTPMethod.Set(span, req.Method)
	if err := d.Tracer.Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header)); err != nil {
		d.Logger.WithError(err).Warn(ctx, "failed injecting trace info")
	}
	return span, clientContext
}

func printContextInternals(ctx interface{}, inner bool) {
	contextValues := reflect.ValueOf(ctx).Elem()
	contextKeys := reflect.TypeOf(ctx).Elem()

	if !inner {
		fmt.Printf("\nFields for %s.%s\n", contextKeys.PkgPath(), contextKeys.Name())
	}

	if contextKeys.Kind() == reflect.Struct {
		for i := 0; i < contextValues.NumField(); i++ {
			reflectValue := contextValues.Field(i)
			reflectValue = reflect.NewAt(reflectValue.Type(), unsafe.Pointer(reflectValue.UnsafeAddr())).Elem()

			reflectField := contextKeys.Field(i)

			if reflectField.Name == "Context" {
				printContextInternals(reflectValue.Interface(), true)
			} else {
				fmt.Printf("field name: %+v\n", reflectField.Name)
				fmt.Printf("value: %+v\n", reflectValue.Interface())
			}
		}
	} else {
		fmt.Printf("context is empty (int)\n")
	}
}
