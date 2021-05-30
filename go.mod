module wedding

go 1.15

require (
	entgo.io/ent v0.8.0
	github.com/XSAM/otelsql v0.3.0
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.7.2
	github.com/go-redis/redis/v8 v8.9.0
	github.com/google/wire v0.5.0
	github.com/lib/pq v1.10.2
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	github.com/toorop/gin-logrus v0.0.0-20210225092905-2c785434f26f
	github.com/urfave/cli/v2 v2.3.0
	go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin v0.20.0
	go.opentelemetry.io/otel v0.20.0
	go.opentelemetry.io/otel/exporters/otlp v0.20.0
	go.opentelemetry.io/otel/metric v0.20.0
	go.opentelemetry.io/otel/sdk v0.20.0
	go.opentelemetry.io/otel/sdk/metric v0.20.0
	go.opentelemetry.io/otel/trace v0.20.0
)

replace (
	entgo.io/ent v0.8.0 => github.com/imle/ent v0.5.1-0.20210530060012-b10bc4779d2c
	github.com/toorop/gin-logrus v0.0.0-20210225092905-2c785434f26f => github.com/imle/gin-logrus v0.0.0-20210402041041-59841976bd0b
)
