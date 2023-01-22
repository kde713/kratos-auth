package core

// Identity represents authenticated identity
type Identity struct {
	// ProviderName is name of authenticated Provider
	ProviderName ProviderName

	// Identifier is unique id of authenticated user/agent
	Identifier string

	// Extra is field for adding custom data
	Extra interface{}
}

// Empty returns true if unauthenticated.
func (i *Identity) Empty() bool {
	return i == nil || i.ProviderName == "" || i.Identifier == ""
}
