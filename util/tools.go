package util

import "errors"

var ErrEmpty = errors.New("data is not exist")

func ToStringSlice(slice []interface{}) []string {
	var result []string
	for _, e := range slice {
		result = append(result, e.(string))
	}
	return result
}

func CheckInterfaceIsString(inter interface{}) bool {
	switch inter.(type) {
	case string:
		return true
	default:
		return false
	}
}
