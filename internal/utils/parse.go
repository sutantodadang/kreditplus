package utils

import "strconv"

func ParseStringToFloat(value string) float64 {

	valFloat, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}

	return valFloat

}

func ParseFloatToString(value float64) string {

	valString := strconv.FormatFloat(value, 'f', -1, 64)

	return valString

}
