package argument_resolver

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

func (resolver *ConversionResolver) Get(name string) (interface{}, bool) {
	return resolver.childResolver.Get(name)
}

func (resolver *ConversionResolver) String(name string) (string, bool) {
	unresolvedData, ok := resolver.Get(name)
	if !ok {
		return "", ok
	}

	return conversion.ToString(unresolvedData)
}

func (resolver *ConversionResolver) Int(name string) (int, bool) {
	unresolvedData, ok := resolver.Get(name)
	if !ok {
		return 0, ok
	}

	return conversion.ToInt(unresolvedData)
}


func (resolver *ConversionResolver) Float(name string) (float64, bool) {
	unresolvedData, ok := resolver.Get(name)
	if !ok {
		return 0, ok
	}

	return conversion.ToFloat(unresolvedData)
}

func (resolver *ConversionResolver) Bool(name string) (bool, bool) {
	unresolvedData, ok := resolver.Get(name)
	if !ok {
		return false, ok
	}

	return conversion.ToBool(unresolvedData)
}
