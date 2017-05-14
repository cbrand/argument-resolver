package conversion

import "strconv"

func ToInt(value interface{}) (int, bool) {
	var data int
	switch converted := value.(type) {
	case string:
		var err error
		data, err = strconv.Atoi(converted)
		if err != nil {
			dataFloat64, err := strconv.ParseFloat(converted, 64)
			if err != nil {
				return 0, false
			}
			data = int(dataFloat64)
		}
	case float64:
		data = int(converted)
	case float32:
		data = int(converted)
	case int:
		data = converted
	default:
		return 0, false
	}
	return data, true
}
