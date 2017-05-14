package resolver


// Resolver defines the interface to request information from an underlying map interface.
type Resolver interface {
	// Get returns and underlying element with and returns whether the element exists. If no extraction could be
	// done the second bool variable will be "false"
	Get(name string) (interface{}, bool)
	// String returns a string if the underlying data could be inferred as such. If no extraction could be done
	// the second bool variable will be "false"
	String(name string) (string, bool)
	// Int returns an int if the underlying data could be inferred as such. If no extraction could be done
	// the second bool variable will be "false"
	Int(name string) (int, bool)
	// Float returns a float object if the underlying data could be inferred as such. If no extraction could be done
	// the second bool variable will be "false"
	Float(name string) (float64, bool)
	// Bool returns a boolean if the underlying data could be inferred as such. If no extraction could be done
	// the second bool variable will be "false"
	Bool(name string) (bool, bool)
}

// DefaultResolver interface provides additional support for resolvers to allow the definition of default values.
type DefaultResolver interface {
	Resolver
	// GetOr returns the specified value or an defaultValue if none could be extracted
	GetOr(name string, defaultValue interface{}) interface{}
	// StringOr returns the specified string. If none could be extracted the default value is returned.
	StringOr(name string, defaultValue string) string
	// IntOr returns the specified int. If none could be extracted the default value is returned.
	IntOr(name string, defaultValue int) int
	// FloatOr returns the specified float. If none could be extracted the default value is returned.
	FloatOr(name string, defaultValue float64) float64
	// BoolOr returns the specified bool.  If none could be extracted the default value is returned.
	BoolOr(name string, defaultValue bool) bool
}
