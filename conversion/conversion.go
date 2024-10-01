package conversion

import (
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	floats := make([]float64, len(strings))

	for index, value := range strings {
		floatVal, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, err
		}
		floats[index] = floatVal
	}

	return floats, nil
}
