package client

import (
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/configuration"
	"bitbucket.org/HeilaSystems/transport/client"
	"bitbucket.org/HeilaSystems/transport/client/http/interceptors/contextValuesToHeaders"
)

func DefaultContextValuesToHeaders(config configuration.Config) client.HTTPClientInterceptor {
	headers , err := config.Get("contextHeaders").StringSlice()
	if err != nil {
		return contextValuesToHeaders.ContextValuesToHeaders([]string{})
	}
	return contextValuesToHeaders.ContextValuesToHeaders(headers)
}

func DefaultTokenClientInterceptors() client.HTTPClientInterceptor {
	return contextValuesToHeaders.TokenClientInterceptors()
}

func DefaultServiceNameToHeader(config configuration.Config)(client.HTTPClientInterceptor,error)  {
	dockerName , err := config.Get("DOCKER_NAME").String()
	if err != nil {
		return nil, err
	}
	return contextValuesToHeaders.ServiceNameToHeader(dockerName),nil
}