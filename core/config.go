package core

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
)

type config struct {
	// useProviders contains ProviderName(s) of registered Provider(s).
	// Authentication order will be same as order of UseProviders.
	useProviders []ProviderName

	// emptyIdentityErrorBuilder is a function that builds kratos errors.Error which will be return when identity is empty on all Provider(s).
	// As the first parameter, kratos request context will be passed.
	emptyIdentityErrorBuilder func(ctx context.Context) *errors.Error
}

type ConfigFunc func(c *config)

var (
	defaultEmptyIdentityError = errors.Unauthorized("unauthorized", "unauthorized")
)

// defaultConfig provides default option for configuring middleware.
func defaultConfig() ConfigFunc {
	return func(c *config) {
		c.emptyIdentityErrorBuilder = func(ctx context.Context) *errors.Error {
			return defaultEmptyIdentityError
		}
		c.useProviders = registerOrder
	}
}

// UseCustomEmptyIdentityError defines custom emptyIdentityError as e
func UseCustomEmptyIdentityError(e *errors.Error) ConfigFunc {
	return func(c *config) {
		c.emptyIdentityErrorBuilder = func(ctx context.Context) *errors.Error {
			return e
		}
	}
}

// UseCustomEmptyIdentityErrorBuilder defines custom emptyIdentityErrorBuilder as f
func UseCustomEmptyIdentityErrorBuilder(f func(ctx context.Context) *errors.Error) ConfigFunc {
	return func(c *config) {
		c.emptyIdentityErrorBuilder = f
	}
}

// AllowEmptyIdentity allows unauthorized request to middleware
func AllowEmptyIdentity() ConfigFunc {
	return func(c *config) {
		c.emptyIdentityErrorBuilder = func(ctx context.Context) *errors.Error {
			return nil
		}
	}
}

// UseProviders defines Provider(s) that will be used in middleware as given ProviderName(s).
func UseProviders(names []ProviderName) ConfigFunc {
	return func(c *config) {
		c.useProviders = names
	}
}
