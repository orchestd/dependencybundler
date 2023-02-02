package client

import (
	"github.com/orchestd/dependencybundler/bundler/contextHeader"
	"github.com/orchestd/dependencybundler/interfaces/configuration"
	"github.com/orchestd/transport/client"
	"github.com/orchestd/transport/client/http/interceptors/contextValuesToHeaders"
)

func DefaultContextValuesToHeaders(config configuration.Config) client.HTTPClientInterceptor {
	headers, err := config.Get("contextHeaders").StringSlice()
	if err != nil {
		return contextValuesToHeaders.ContextValuesToHeaders([]string{})
	}
	headers = append(headers, contextHeader.AlwaysCopyHeaders...)
	return contextValuesToHeaders.ContextValuesToHeaders(headers)
}

func DefaultTokenClientInterceptors() client.HTTPClientInterceptor {
	return contextValuesToHeaders.TokenClientInterceptors()
}

func DefaultServiceNameToHeader(config configuration.Config) (client.HTTPClientInterceptor, error) {
	dockerName, err := config.GetServiceName()
	if err != nil {
		return nil, err
	}
	return contextValuesToHeaders.ServiceNameToHeader(dockerName), nil
}
