package tracer

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/uber/jaeger-client-go"

	"github.com/opentracing/opentracing-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func InitJaeger(service string) {
	cfg, _ := jaegercfg.FromEnv()
	cfg.ServiceName = service
	tracer, _, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))

	if err != nil {
		klog.Fatalf("Jaeger initialization failed: %v", err)
	}

	opentracing.InitGlobalTracer(tracer)
}
