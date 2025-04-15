package conversion

import (
	"errors"
	"strconv"
)

type in []string
type out []float64

func StringsToFloats(values in) (out, error) {
	out := make([]float64, len(values))

	for i, val := range values {
		newVal, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return nil, errors.New("error at string to float conversion")
		}
		out[i] = newVal
	}

	return out, nil
}
