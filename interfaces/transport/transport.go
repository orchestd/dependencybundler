package transport

import (
	"bitbucket.org/HeilaSystems/transport/client"
	clientHTTP "bitbucket.org/HeilaSystems/transport/client/http"
	"bitbucket.org/HeilaSystems/transport/discoveryService"
	"bitbucket.org/HeilaSystems/transport/server"
	"bitbucket.org/HeilaSystems/transport/server/http"
	"github.com/gin-gonic/gin"
)

type IRouter gin.IRouter
type IHandler server.IHandler
type Handler server.Handler

var NewHttpHandler = server.NewHttpHandler

var HandleFunc = http.HandleFunc

type HttpClient client.HttpClient
type HTTPClientInterceptor client.HTTPClientInterceptor
type IHttpLog http.IHttpLog
type DiscoveryServiceProvider discoveryService.DiscoveryServiceProvider
type HttpServerSettings http.HttpServerSettings

const ContentTypeJSON = clientHTTP.ContentTypeJSON
const ContentTypeXML = clientHTTP.ContentTypeXML
