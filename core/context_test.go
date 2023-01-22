package core

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func Test_ExtractEmptyIdentity(t *testing.T) {
	identity := IdentityFromContext(context.TODO())
	if !identity.Empty() {
		t.Errorf("identity is not empty: %#v", identity)
	}
}

func Test_SetAndGetIdentity(t *testing.T) {
	testIdentity := Identity{
		ProviderName: "test_provider",
		Identifier:   "test_id",
		Extra: struct {
			Foo string
			Bar time.Time
		}{
			Foo: "foo",
			Bar: time.Now(),
		},
	}
	ctx := context.WithValue(context.Background(), "test", "foo.bar")
	identityCtx := contextWithIdentity(ctx, testIdentity)
	retrievedIdentity := IdentityFromContext(identityCtx)
	if diff := cmp.Diff(testIdentity, retrievedIdentity); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}
