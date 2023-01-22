package core

var registerOrder []ProviderName
var registeredProviders map[ProviderName]Provider

func init() {
	registeredProviders = make(map[ProviderName]Provider)
}

// RegisterProvider registers Provider with ProviderName
func RegisterProvider(name ProviderName, provider Provider) {
	registeredProviders[name] = provider
	registerOrder = append(registerOrder, name)
}
