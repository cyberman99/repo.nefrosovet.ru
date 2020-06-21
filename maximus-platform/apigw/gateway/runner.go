package gateway

import (
	"context"

	"github.com/sirupsen/logrus"

	krakendbf "github.com/devopsfaith/bloomfilter/krakend"
	"github.com/devopsfaith/krakend-jose"
	krakendlogrus "github.com/devopsfaith/krakend-logrus"
	_ "github.com/devopsfaith/krakend-opencensus/exporter/influxdb"
	_ "github.com/devopsfaith/krakend-opencensus/exporter/jaeger"
	_ "github.com/devopsfaith/krakend-opencensus/exporter/prometheus"
	_ "github.com/devopsfaith/krakend-opencensus/exporter/stackdriver"
	_ "github.com/devopsfaith/krakend-opencensus/exporter/xray"
	_ "github.com/devopsfaith/krakend-opencensus/exporter/zipkin"
	"github.com/devopsfaith/krakend/config"
	krakendrouter "github.com/devopsfaith/krakend/router"
	krakendhttptreemux "github.com/devopsfaith/krakend/router/httptreemux"
	"github.com/devopsfaith/krakend/router/mux"
	"github.com/dimfeld/httptreemux"
)

func Run(ctx context.Context, cfg config.ServiceConfig) {
	logger := krakendlogrus.WrapLogger(logrus.StandardLogger(), "KrakenD")

	logger.Info("Listening on port:", cfg.Port)

	rejecter, err := krakendbf.Register(ctx, "krakend-bf", cfg, logger, func(n string, p int) {})
	if err != nil {
		logger.Warning("bloomFilter:", err.Error())
	}

	// setup the krakend router
	routerFactory := mux.NewFactory(mux.Config{
		Engine:         krakendhttptreemux.NewEngine(httptreemux.NewContextMux()),
		ProxyFactory:   NewProxyFactory(logger, NewBackendFactory(logger)),
		Middlewares:    []mux.HandlerMiddleware{},
		Logger:         logger,
		HandlerFactory: NewHandlerFactory(logger, jose.RejecterFunc(rejecter.RejectToken)),
		RunServer:      krakendrouter.RunServer,
	})

	// start the engines
	routerFactory.NewWithContext(ctx).Run(cfg)
}
