package common

import (
	"strconv"
)

func ResultToString(data interface{}) string{
    return string(data.([]uint8))
}

func IntToString(data int) string{
    return strconv.Itoa(data)
}

func StringToInt(data string) int{
    num, _ := strconv.Atoi(data)

    return num
}

func FloatToString(data float64) string{
    return strconv.FormatFloat(data, 'f', -1, 32)
}

func BoolToInt(data bool) (result int){
    result = 0

    if data{
        result = 1
    }

    return
}