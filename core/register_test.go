package core

import (
	"context"
	"testing"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

var _ Provider = (*testProvider)(nil)

type testProvider struct {
	RawName string
}

func (t testProvider) Name() ProviderName {
	return ProviderName(t.RawName)
}

func (t testProvider) Authenticate(_ context.Context, _ interface{}) (*Identity, *errors.Error) {
	return nil, nil
}

func Test_RegisterProvider(t *testing.T) {
	RegisterProvider("test1", testProvider{"test1"})
	RegisterProvider("test2", testProvider{"test2"})
	expectProvidersMap := map[ProviderName]Provider{
		"test1": testProvider{"test1"},
		"test2": testProvider{"test2"},
	}
	if diff := cmp.Diff(expectProvidersMap, registeredProviders, cmpopts.IgnoreUnexported()); diff != "" {
		t.Errorf("registeredProviders mismatch (-want +got):\n%s", diff)
	}
	expectProvidersOrder := []ProviderName{"test1", "test2"}
	if diff := cmp.Diff(expectProvidersOrder, registerOrder); diff != "" {
		t.Errorf("registerOrder mismatch (-want +got):\n%s", diff)
	}
}
