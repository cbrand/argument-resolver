package conversion

import "strconv"

// ToFloat tries to convert a passed value to float64. The second value is set to false if no conversion could be done.
func ToFloat(value interface{}) (float64, bool) {
	var data float64
	switch converted := value.(type) {
	case string:
		var err error
		data, err = strconv.ParseFloat(converted, 64)
		if err != nil {
			return 0.0, false
		}
	case float64:
		data = converted
	case float32:
		data = float64(converted)
	case int:
		data = float64(converted)
	default:
		return 0.0, false
	}
	return data, true
}
