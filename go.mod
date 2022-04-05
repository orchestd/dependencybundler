module bitbucket.org/HeilaSystems/dependencybundler

go 1.14

require (
	bitbucket.org/HeilaSystems/cacheStorage v0.16.2
	bitbucket.org/HeilaSystems/configurations v0.3.0
	bitbucket.org/HeilaSystems/debug v0.0.1
	bitbucket.org/HeilaSystems/log v0.0.11
	bitbucket.org/HeilaSystems/monitoring v0.1.0
	bitbucket.org/HeilaSystems/servicereply v0.0.3
	bitbucket.org/HeilaSystems/session v0.16.1
	bitbucket.org/HeilaSystems/sharedlib v0.3.0
	bitbucket.org/HeilaSystems/trace v0.0.10
	bitbucket.org/HeilaSystems/transport v0.11.1
	bitbucket.org/HeilaSystems/validations v0.3.0
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.7.2
	github.com/go-masonry/mortar v0.1.3
	github.com/opentracing/opentracing-go v1.2.0
	github.com/prometheus/client_golang v1.11.0
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	go.uber.org/fx v1.13.1
	google.golang.org/api v0.30.0
	google.golang.org/grpc v1.34.0
)

replace (
	bitbucket.org/HeilaSystems/cacheStorage v0.16.2 => ../cacheStorage
    bitbucket.org/HeilaSystems/configurations v0.3.0 => ../configurations
    bitbucket.org/HeilaSystems/log v0.0.11 => ../log
    bitbucket.org/HeilaSystems/monitoring v0.1.0 => ../monitoring
    bitbucket.org/HeilaSystems/session v0.16.1 => ../session
    bitbucket.org/HeilaSystems/transport v0.11.1 => ../transport
)
