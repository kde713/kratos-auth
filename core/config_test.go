package core

import (
	"testing"

	"github.com/go-kratos/kratos/v2/errors"
)

func Test_ConfigFunc(t *testing.T) {
	testCases := []struct {
		name    string
		funcs   []ConfigFunc
		matcher func(c *config) bool
	}{
		{
			name: "use_custom_error",
			funcs: []ConfigFunc{
				UseCustomEmptyIdentityError(errors.BadRequest("test", "test")),
			},
			matcher: func(c *config) bool {
				return errors.IsBadRequest(c.emptyIdentityError)
			},
		},
		{
			name: "allow_empty_identity",
			funcs: []ConfigFunc{
				AllowEmptyIdentity(),
			},
			matcher: func(c *config) bool {
				return c.emptyIdentityError == nil
			},
		},
		{
			name: "use_providers",
			funcs: []ConfigFunc{
				UseProviders([]ProviderName{"test1", "test2"}),
			},
			matcher: func(c *config) bool {
				return c.useProviders[0] == "test1" && c.useProviders[1] == "test2"
			},
		},
		{
			name: "complex",
			funcs: []ConfigFunc{
				UseProviders([]ProviderName{"test"}),
				UseCustomEmptyIdentityError(errors.Forbidden("test", "test")),
			},
			matcher: func(c *config) bool {
				return errors.IsForbidden(c.emptyIdentityError) && c.useProviders[0] == "test"
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var c config
			defaultConfig()(&c)
			for _, f := range tc.funcs {
				f(&c)
			}
			if !tc.matcher(&c) {
				t.Errorf("failed to match: %#v", c)
			}
		})
	}
}
