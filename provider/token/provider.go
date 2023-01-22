package token

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport"

	"github.com/kde713/kratos-auth/core"
)

var _ core.Provider = (*provider)(nil)

type provider struct {
	name      core.ProviderName
	scheme    string
	validator func(token string) (string, interface{}, *errors.Error)
}

func (p provider) Name() core.ProviderName {
	return p.name
}

func (p provider) Authenticate(ctx context.Context, _ interface{}) (*core.Identity, *errors.Error) {
	var token string
	tr, ok := transport.FromServerContext(ctx)
	if ok {
		authHeader := tr.RequestHeader().Get("authorization")
		authHeaderParts := strings.Split(authHeader, fmt.Sprintf("%s ", p.scheme))
		if len(authHeaderParts) == 2 && authHeaderParts[0] == "" {
			token = authHeaderParts[1]
		}
	}
	id, extra, httpErr := p.validator(token)
	return &core.Identity{
		ProviderName: p.Name(),
		Identifier:   id,
		Extra:        extra,
	}, httpErr
}

func Register(
	name core.ProviderName,
	scheme string,
	validator func(token string) (string, interface{}, *errors.Error),
) {
	core.RegisterProvider(
		name,
		provider{
			name:      name,
			scheme:    scheme,
			validator: validator,
		},
	)
}
