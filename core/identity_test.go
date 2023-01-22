package core_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/kde713/kratos-auth/core"
)

func Test_IdentityEmpty(t *testing.T) {
	testCases := []struct {
		name     string
		identity core.Identity
		isEmpty  bool
	}{
		{
			name:     "full_empty",
			identity: core.Identity{},
			isEmpty:  true,
		},
		{
			name: "empty_provider",
			identity: core.Identity{
				Identifier: "test_id",
				Extra:      struct{ CreateTime time.Time }{time.Now()},
			},
			isEmpty: true,
		},
		{
			name: "empty_identifier",
			identity: core.Identity{
				ProviderName: "test_provider",
				Extra:        struct{ CreateTime time.Time }{time.Now()},
			},
			isEmpty: true,
		},
		{
			name: "valid",
			identity: core.Identity{
				ProviderName: "test_provider",
				Identifier:   "test_id",
			},
			isEmpty: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if diff := cmp.Diff(tc.isEmpty, tc.identity.Empty()); diff != "" {
				t.Errorf(".Empty() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
