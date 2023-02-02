package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/orchestd/transport/client"
	clientHTTP "github.com/orchestd/transport/client/http"
	"github.com/orchestd/transport/discoveryService"
	"github.com/orchestd/transport/server"
	"github.com/orchestd/transport/server/http"
)

type IRouter gin.IRouter
type IHandler server.IHandler
type Handler server.Handler

var NewHttpHandler = server.NewHttpHandler

var HandleFunc = http.HandleFunc
var FileReplyHandleFunc = http.FileReplyHandleFunc

type HttpClient client.HttpClient
type HTTPClientInterceptor client.HTTPClientInterceptor
type IHttpLog http.IHttpLog
type DiscoveryServiceProvider discoveryService.DiscoveryServiceProvider
type HttpServerSettings http.HttpServerSettings

const ContentTypeJSON = clientHTTP.ContentTypeJSON
const ContentTypeXML = clientHTTP.ContentTypeXML
