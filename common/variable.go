package common

func ToString(data interface{}) string{
    return string(data.([]uint8))
}