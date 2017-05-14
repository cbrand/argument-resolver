package resolver

// NewArgumentResolver returns a resolver which does not try to do any type conversions before returning the
// object.
func NewArgumentResolver(data map[string]interface{}) DefaultResolver {
	argumentResolver := &ArgumentResolver{
		data: data,
	}
	return NewDefaultResolver(argumentResolver)
}

// ArgumentResolver implements the resolver interface and limits the conversion to direct translations
// and does not try to convert data between formats.
type ArgumentResolver struct {
	data map[string]interface{}
}

// Get returns and underlying element with and returns whether the element exists
func (resolver *ArgumentResolver) Get(name string) (interface{}, bool) {
	data, ok := resolver.data[name]
	return data, ok
}

// String returns a string if the underlying data could be inferred as such
func (resolver *ArgumentResolver) String(name string) (string, bool) {
	untypedData, ok := resolver.Get(name)
	if !ok {
		return "", ok
	}

	data, ok := untypedData.(string)
	return data, ok
}

// Int returns an int if the underlying data could be inferred as such
func (resolver *ArgumentResolver) Int(name string) (int, bool) {
	untypedData, ok := resolver.Get(name)
	if !ok {
		return 0, ok
	}

	data, ok := untypedData.(int)
	return data, ok
}

// Float returns a float object if the underlying data could be inferred as such
func (resolver *ArgumentResolver) Float(name string) (float64, bool) {
	untypedData, ok := resolver.Get(name)
	if !ok {
		return 0, ok
	}

	data, ok := untypedData.(float64)
	if !ok {
		var data32 float32
		data32, ok = untypedData.(float32)
		data = float64(data32)
	}

	return data, ok
}

// Bool returns a boolean if the underlying data could be inferred as such.
func (resolver *ArgumentResolver) Bool(name string) (bool, bool) {
	untypedData, ok := resolver.Get(name)
	if !ok {
		return false, ok
	}

	data, ok := untypedData.(bool)
	return data, ok
}
