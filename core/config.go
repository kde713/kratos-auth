package core

import (
	"github.com/go-kratos/kratos/v2/errors"
)

type config struct {
	// useProviders contains ProviderName(s) of registered Provider(s).
	// Authentication order will be same as order of UseProviders.
	useProviders []ProviderName

	// emptyIdentityError is kratos errors.Error which will be return when identity is empty on all Provider(s).
	// If nil, middleware will allow empty identity, and pass through requests.
	emptyIdentityError *errors.Error
}

type ConfigFunc func(c *config)

var (
	defaultEmptyIdentityError = errors.Unauthorized("unauthorized", "unauthorized")
)

// defaultConfig provides default option for configuring middleware.
func defaultConfig() ConfigFunc {
	return func(c *config) {
		c.emptyIdentityError = defaultEmptyIdentityError
		c.useProviders = registerOrder
	}
}

// UseCustomEmptyIdentityError defines custom emptyIdentityError as e
func UseCustomEmptyIdentityError(e *errors.Error) ConfigFunc {
	return func(c *config) {
		c.emptyIdentityError = e
	}
}

// AllowEmptyIdentity allows unauthorized request to middleware
func AllowEmptyIdentity() ConfigFunc {
	return func(c *config) {
		c.emptyIdentityError = nil
	}
}

// UseProviders defines Provider(s) that will be used in middleware as given ProviderName(s).
func UseProviders(names []ProviderName) ConfigFunc {
	return func(c *config) {
		c.useProviders = names
	}
}
