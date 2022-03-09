package convert

import (
	"strconv"
)

func ResultToString(data interface{}) string {
	return string(data.([]uint8))
}

func IntToString(data int) string {
	return strconv.Itoa(data)
}

func StringToInt(data string) int {
	num, _ := strconv.Atoi(data)

	return num
}

func FloatToString(data float64) string {
	return strconv.FormatFloat(data, 'f', -1, 32)
}

func StringToFloat(data string) float64 {
	num, _ := strconv.ParseFloat(data, 64)
	return num
}

func BoolToInt(data bool) (result int) {
	result = 0

	if data {
		result = 1
	}

	return
}

func BoolToString(data bool) (result string) {
	return IntToString(BoolToInt(data))
}

func StringToBool(data string) (result bool) {
	result = true
	if data == "0" || data == "false" {
		result = false
	}
	return
}
