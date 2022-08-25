package utils

func ConvertNilToZeroValue[T any](value *T) T {
	if value == nil {
		var result T
		return result
	}
	return *value
}
