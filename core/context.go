package core

import (
	"context"
)

// identityContextKey is key which Identity stored in context.Context
type identityContextKey struct{}

// contextWithIdentity returns new context.Context with given Identity
func contextWithIdentity(ctx context.Context, i Identity) context.Context {
	if i.Empty() {
		return ctx
	}
	return context.WithValue(ctx, identityContextKey{}, i)
}

// IdentityFromContext returns Identity which stored in context.Context.
// NOTE: Identity can be empty depends on config. If you defined emptyIdentityError as nil, you must check emptiness with Identity.Empty()
func IdentityFromContext(ctx context.Context) Identity {
	i, _ := ctx.Value(identityContextKey{}).(Identity)
	return i
}
