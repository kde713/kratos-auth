package core

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
)

// ProviderName is unique name of provider
type ProviderName string

type Provider interface {
	// Name returns registered ProviderName
	Name() ProviderName

	// Authenticate extracts Identity from server context and request
	// If you return non-nil error as second parameter, middleware will respond error directly.
	// Error must be errors.Error of kratos
	Authenticate(ctx context.Context, req interface{}) (*Identity, *errors.Error)
}
