package gateway

import (
	"github.com/devopsfaith/krakend-jose"
	"github.com/devopsfaith/krakend/logging"
	"github.com/devopsfaith/krakend/router/httptreemux"
	"github.com/devopsfaith/krakend/router/mux"
	"repo.nefrosovet.ru/maximus-platform/apigw/gateway/validator"
)

// NewHandlerFactory returns a HandlerFactory
func NewHandlerFactory(logger logging.Logger, rejecter jose.Rejecter) mux.HandlerFactory {
	handlerFactory := mux.CustomEndpointHandler(mux.NewRequestBuilder(httptreemux.ParamsExtractor))
	handlerFactory = validator.HandlerFactory(handlerFactory, httptreemux.ParamsExtractor, logger, rejecter)

	return handlerFactory
}
