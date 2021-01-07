package transport

import (
	"bitbucket.org/HeilaSystems/transport/client"
	"bitbucket.org/HeilaSystems/transport/server/http"
	"github.com/gin-gonic/gin"
)


type IRouter gin.IRouter
var HandleFunc =  http.HandleFunc

type HttpClient client.HttpClient
type HTTPClientInterceptor client.HTTPClientInterceptor
type IHttpLog http.IHttpLog