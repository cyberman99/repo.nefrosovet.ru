package gateway

import (
	"github.com/devopsfaith/krakend-jsonschema"
	"github.com/devopsfaith/krakend/logging"
	"github.com/devopsfaith/krakend/proxy"
)

// NewProxyFactory returns a new ProxyFactory wrapping the injected BackendFactory with the default proxy stack and a metrics collector
func NewProxyFactory(logger logging.Logger, backendFactory proxy.BackendFactory) proxy.Factory {
	proxyFactory := proxy.NewDefaultFactory(backendFactory, logger)
	proxyFactory = jsonschema.ProxyFactory(proxyFactory)
	// proxyFactory = metricCollector.ProxyFactory("pipe", proxyFactory)
	// proxyFactory = opencensus.ProxyFactory(proxyFactory)
	return proxyFactory
}
