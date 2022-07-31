package providers

import (
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/configuration"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/transport"
	"bitbucket.org/HeilaSystems/log"
	"bitbucket.org/HeilaSystems/servicereply"
	"bitbucket.org/HeilaSystems/sharedlib/consts"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
)

type simpleHttpDSP struct {
	serviceName string
	port        string
	client      transport.HttpClient
	conf        configuration.Config
	logger      log.Logger
	server      transport.HttpServerSettings
}

const discoveryServiceName = "discoveryService"
const discoveryServiceAddress = "http://127.0.0.1:8500" //TODO : env
const address = "address"

func NewSimpleHttpDSP(client transport.HttpClient, conf configuration.Config, logger log.Logger) simpleHttpDSP { //serviceName string,port string
	errBase := "NewSimpleHttpDSP: "

	s := simpleHttpDSP{client: client, logger: logger}

	if confPort, err := conf.Get("port").String(); err != nil {
		logger.WithError(err).Info(context.Background(), errBase+"Cannot get port from configuration")
		//should panic ?
	} else {
		s.port = confPort
	}

	if dn, err := conf.Get(consts.ServiceNameEnv).String(); err != nil {
		logger.WithError(err).Debug(context.Background(), errBase+"Cannot get "+consts.ServiceNameEnv+" from configurations")
		//should panic ?
	} else {
		s.serviceName = dn
	}

	return s
}

func (dsp simpleHttpDSP) Register() (sr servicereply.ServiceReply) {
	req := gin.H{"serviceName": dsp.serviceName, "method": "isAlive", "port": dsp.port}
	var res interface{}

	sRep := dsp.client.Post(context.Background(), req, discoveryServiceName, "register", &res, nil)

	return sRep.WithError(fmt.Errorf("cant register service %s using port %s in simple discovery service", dsp.serviceName, dsp.port))
}

func (dsp simpleHttpDSP) GetAddress(serviceName string) (sr servicereply.ServiceReply) {
	var retVal string
	if serviceName == discoveryServiceName {
		retVal = discoveryServiceAddress
	} else {
		req := gin.H{"serviceName": serviceName}
		var res struct {
			Address string `json:"address"`
		}

		if sRep := dsp.client.Call(context.Background(), req, discoveryServiceName, "getAddress", &res, nil); !sRep.IsSuccess() {
			return sRep
		} else {
			if res.Address == "" {
				servicereply.NewNoMatchReply("reply from discovery service for " + serviceName + " does not include address value")
			} else {
				retVal = res.Address
			}

		}

	}
	return servicereply.NewNil().WithReplyValues(servicereply.ValuesMap{address: retVal})
}
