package core

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
)

func NewMiddleware(configFuncs ...ConfigFunc) middleware.Middleware {
	var conf config
	defaultConfig()(&conf)
	for _, configFunc := range configFuncs {
		configFunc(&conf)
	}

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			var identity Identity
			for _, providerName := range conf.useProviders {
				provider, ok := registeredProviders[providerName]
				if !ok {
					continue
				}
				i, e := provider.Authenticate(ctx, req)
				if e != nil {
					return nil, e
				}
				if i != nil {
					identity = *i
					break
				}
			}
			if identity.Empty() {
				emptyIdentityErr := conf.emptyIdentityErrorBuilder(ctx)
				if emptyIdentityErr != nil {
					return nil, emptyIdentityErr
				}
			}
			return handler(contextWithIdentity(ctx, identity), req)
		}
	}
}
