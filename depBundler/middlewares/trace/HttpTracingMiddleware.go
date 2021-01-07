package trace

import (
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/configuration"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/log"
	"bytes"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-masonry/mortar/utils"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	traceLog "github.com/opentracing/opentracing-go/log"
	"go.uber.org/fx"
	"google.golang.org/grpc/metadata"
	"io"
	"io/ioutil"
)

type tracingDeps struct {
	fx.In

	Logger log.Logger
	Config configuration.Config
	Tracer opentracing.Tracer `optional:"true"`
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

var grpcTag = opentracing.Tag{Key: string(ext.Component), Value: "gRPC"}
var restTag = opentracing.Tag{Key: string(ext.Component), Value: "REST"}

func HttpTracingUnaryServerInterceptor(deps tracingDeps) gin.HandlerFunc {
	return func(c *gin.Context) {

		if deps.Tracer == nil {
			c.Next()
			return
		}
		deps.Logger.Debug(c,"hello")
		span, ctx := deps.newServerSpan(c, c.Request.RequestURI)
		c.Request.WithContext(ctx)
		defer span.Finish()

		bodyCopy := new(bytes.Buffer)
		io.Copy(bodyCopy, c.Request.Body)
		bodyData := bodyCopy.Bytes()
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(bodyData))
		blw := &bodyLogWriter{body: bytes.NewBuffer([]byte{}), ResponseWriter: c.Writer}
		c.Writer = blw
		reqJson := json.RawMessage(bodyData)
		addBodyToSpan(span, "request", reqJson)
		// call handler
		c.Next()
		rawBody := blw.body
		jsonmsg := json.RawMessage(rawBody.Bytes())
		if len(c.Errors.String()) > 0 {
			ext.LogError(span, c.Errors[0].Err)
		} else {
			addBodyToSpan(span, "response", jsonmsg)
		}
	}
}

func (d tracingDeps) newServerSpan(ctx context.Context, methodName string) (opentracing.Span, context.Context) {
	spanContext, extractError := d.Tracer.Extract(opentracing.HTTPHeaders, d.extractIncomingCarrier(ctx))
	if extractError != nil && extractError != opentracing.ErrSpanContextNotFound {
		d.Logger.WithError(extractError).Debug(ctx, "failed extracting trace info") // really low level information in my opinion
	}
	return opentracing.StartSpanFromContextWithTracer(ctx, d.Tracer, methodName, ext.RPCServerOption(spanContext), ext.SpanKindRPCServer, grpcTag)
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
