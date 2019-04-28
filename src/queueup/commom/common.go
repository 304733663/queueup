package commom

import "fmt"

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func CheckValueEmpty(value interface{}) bool {
	if value == nil {
		return true
	}
	valueType := fmt.Sprintf("%T", value)
	if valueType == "string" && value == "" {
		return true
	}
	return false
}
