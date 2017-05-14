package conversion

import "strconv"

// StringRepresentator is an interface which can be converted via the ToString method to a string representation
type StringRepresentator interface {
	String() string
}

// ToString tries to convert an arbitrary data format into a string representation. The second value is set to false
// if no conversion could be achieved.
func ToString(data interface{}) (string, bool) {
	var convertedData string
	switch element := data.(type) {
	case string:
		convertedData = element
	case int:
		convertedData = strconv.Itoa(element)
	case float64:
		convertedData = strconv.FormatFloat(element, 'f', -1, 64)
	case float32:
		convertedData = strconv.FormatFloat(float64(element), 'f', -1, 32)
	case bool:
		convertedData = strconv.FormatBool(element)
	case StringRepresentator:
		convertedData = element.String()
	default:
		return "", false
	}
	return convertedData, true
}
