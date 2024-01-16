package trace

import (
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-masonry/mortar/utils"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	traceLog "github.com/opentracing/opentracing-go/log"
	"github.com/orchestd/dependencybundler/interfaces/configuration"
	"github.com/orchestd/dependencybundler/interfaces/log"
	"github.com/orchestd/tokenauth"
	"go.uber.org/fx"
	"google.golang.org/grpc/metadata"
)

type tracingDeps struct {
	fx.In
	JWToken tokenauth.TokenBase
	Logger  log.Logger
	Config  configuration.Config
	Tracer  opentracing.Tracer `optional:"true"`
}

func addBodyToSpan(span opentracing.Span, name string, msg interface{}) {
	bytes, err := utils.MarshalMessageBody(msg)
	if err == nil {
		span.LogFields(traceLog.String(name, string(bytes))) // TODO: can exceed length limit, introduce option
	} else {
		// If marshaling failed let's try to log msg.ToString()
		span.LogKV(name, msg)
	}
}

var httpTag = opentracing.Tag{Key: string(ext.Component), Value: "HTTP"}

func (d tracingDeps) newServerSpan(ctx *gin.Context, methodName string) (opentracing.Span, context.Context) {
	spanContext, extractError := d.Tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(ctx.Request.Header))
	if extractError != nil && extractError != opentracing.ErrSpanContextNotFound {
		d.Logger.WithError(extractError).Debug(ctx.Request.Context(), "failed extracting trace info") // really low level information in my opinion
	}
	return opentracing.StartSpanFromContextWithTracer(ctx.Request.Context(), d.Tracer, methodName, ext.RPCServerOption(spanContext), ext.SpanKindRPCServer, httpTag)
}

func (d tracingDeps) extractIncomingCarrier(ctx context.Context) utils.MDTraceCarrier {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	}
	return utils.MDTraceCarrier(md.Copy()) // make a copy since this map is not thread safe
}

func (d tracingDeps) extractOutgoingCarrier(ctx context.Context) utils.MDTraceCarrier {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	}
	return utils.MDTraceCarrier(md.Copy()) // make a copy since this map is not thread safe
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
