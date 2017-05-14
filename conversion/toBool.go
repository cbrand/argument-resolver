package conversion

import (
	"strconv"
	"strings"
)

// ToBool tries to convert a value to bool. If no conversion is possible, the second return value is set to "false".
func ToBool(value interface{}) (bool, bool) {
	var data bool
	switch converted := value.(type) {
	case string:
		var err error
		data, err = strconv.ParseBool(strings.ToLower(converted))
		if err != nil {
			return false, false
		}
	case float64, float32, int:
		data = converted != 0
	case bool:
		data = converted
	default:
		return false, false
	}
	return data, true
}
