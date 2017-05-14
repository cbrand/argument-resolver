package argument_resolver

// NewDefaultResolver wraps a resolver object and adds the necessary functions to create a fitting
// object which implements the DefaultResolver interface.
func NewDefaultResolver(resolver Resolver) DefaultResolver {
	return &DefaultResolverImpl{Resolver: resolver}
}

// DefaultResolverImpl wraps a resolver object and adds the necessary functions to adher to the
// DefaultResolver interface
type DefaultResolverImpl struct {
	Resolver
}

// GetOr returns the specified value or an defaultValue if none could be extracted
func (resolver *DefaultResolverImpl) GetOr(name string, defaultValue interface{}) interface{} {
	data, ok := resolver.Get(name)
	if !ok {
		return defaultValue
	}
	return data
}

// StringOr returns the specified string. If none could be extracted the default value is returned.
func (resolver *DefaultResolverImpl) StringOr(name string, defaultValue string) string {
	data, ok := resolver.String(name)
	if !ok {
		return defaultValue
	}
	return data
}

// IntOr returns the specified int. If none could be extracted the default value is returned.
func (resolver *DefaultResolverImpl) IntOr(name string, defaultValue int) int {
	data, ok := resolver.Int(name)
	if !ok {
		return defaultValue
	}
	return data
}

// FloatOr returns the specified float. If none could be extracted the default value is returned.
func (resolver *DefaultResolverImpl) FloatOr(name string, defaultValue float64) float64 {
	data, ok := resolver.Float(name)
	if !ok {
		return defaultValue
	}
	return data
}

// BoolOr returns the specified bool.  If none could be extracted the default value is returned.
func (resolver *DefaultResolverImpl) BoolOr(name string, defaultValue bool) bool {
	data, ok := resolver.Bool(name)
	if !ok {
		return defaultValue
	}
	return data
}
