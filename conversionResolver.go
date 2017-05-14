package resolver

import "github.com/cbrand/argument-resolver/conversion"

// NewConversionResolver returns a resolver which does try to convert the values before returning them.
// This means that a string "true" will also result into a valid response. Also the bool "true" will be
// converted to the string "true" on return.
func NewConversionResolver(data map[string]interface{}) DefaultResolver {
	argumentResolver := NewArgumentResolver(data)
	conversionResolver := &ConversionResolver{
		childResolver: argumentResolver,
	}
	return NewDefaultResolver(conversionResolver)
}

// ConversionResolver returns a resolver which tries to automatically convert to the wished interface.
type ConversionResolver struct {
	childResolver Resolver
}

// Get returns and underlying element with and returns whether the element exists. If no extraction could be
// done the second bool variable will be "false"
func (resolver *ConversionResolver) Get(name string) (interface{}, bool) {
	return resolver.childResolver.Get(name)
}

// String returns a string if the underlying data could be inferred as such. If no extraction could be done
// the second bool variable will be "false"
func (resolver *ConversionResolver) String(name string) (string, bool) {
	unresolvedData, ok := resolver.Get(name)
	if !ok {
		return "", ok
	}

	return conversion.ToString(unresolvedData)
}

// Int returns an int if the underlying data could be inferred as such. If no extraction could be done
// the second bool variable will be "false"
func (resolver *ConversionResolver) Int(name string) (int, bool) {
	unresolvedData, ok := resolver.Get(name)
	if !ok {
		return 0, ok
	}

	return conversion.ToInt(unresolvedData)
}


// Float returns a float object if the underlying data could be inferred as such. If no extraction could be done
// the second bool variable will be "false"
func (resolver *ConversionResolver) Float(name string) (float64, bool) {
	unresolvedData, ok := resolver.Get(name)
	if !ok {
		return 0, ok
	}

	return conversion.ToFloat(unresolvedData)
}

// Bool returns a boolean if the underlying data could be inferred as such. If no extraction could be done
// the second bool variable will be "false"
func (resolver *ConversionResolver) Bool(name string) (bool, bool) {
	unresolvedData, ok := resolver.Get(name)
	if !ok {
		return false, ok
	}

	return conversion.ToBool(unresolvedData)
}
