module github.com/orchestd/dependencybundler

go 1.18

require github.com/orchestd/nats.io v0.0.0-20231212092442-e8d6ed802a1f

replace github.com/orchestd/nats.io v0.0.0-20231212092442-e8d6ed802a1f => ../../../github.com/orchestd/nats.io

require (
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.8.1
	github.com/go-masonry/mortar v0.1.3
	github.com/opentracing/opentracing-go v1.2.0
	github.com/orchestd/cacheStorage v0.22.0
	github.com/orchestd/configurations v0.13.1
	github.com/orchestd/debug v0.1.11
	github.com/orchestd/log v0.1.3
	github.com/orchestd/monitoring v0.2.3
	github.com/orchestd/servicereply v0.0.8
	github.com/orchestd/session v0.22.0
	github.com/orchestd/sharedlib v0.16.0
	github.com/orchestd/tokenauth v0.4.15
	github.com/orchestd/trace v0.0.23
	github.com/orchestd/transport v0.16.1
	github.com/orchestd/validations v0.5.12
	github.com/prometheus/client_golang v1.11.0
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	go.uber.org/fx v1.18.1
	google.golang.org/api v0.30.0
	google.golang.org/grpc v1.34.0
)

require (
	cloud.google.com/go v0.65.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.11.0 // indirect
	github.com/goccy/go-json v0.9.11 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/googleapis/gax-go/v2 v2.0.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kelseyhightower/envconfig v1.4.0 // indirect
	github.com/klauspost/compress v1.17.0 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe // indirect
	github.com/nats-io/nats.go v1.31.0 // indirect
	github.com/nats-io/nkeys v0.4.5 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/orchestd/serviceerror v0.4.2 // indirect
	github.com/orchestd/utils v0.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.3 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.26.0 // indirect
	github.com/prometheus/procfs v0.6.0 // indirect
	github.com/rs/zerolog v1.20.0 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	go.mongodb.org/mongo-driver v1.12.1 // indirect
	go.opencensus.io v0.22.4 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/dig v1.15.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.22.0 // indirect
	golang.org/x/crypto v0.6.0 // indirect
	golang.org/x/net v0.6.0 // indirect
	golang.org/x/oauth2 v0.0.0-20200902213428-5d25da1a8d43 // indirect
	golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/genproto v0.0.0-20201007142714-5c0e72c5e71e // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
