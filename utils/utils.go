package utils

import "strings"

func ConvertNilToZeroValue[T any](value *T) T {
	if value == nil {
		var result T
		return result
	}
	return *value
}

func GetQuestionMarkStrs(num int) string {
	return strings.Join(strings.Split(strings.Repeat("?", num), ""), ",")
}

func IgnoreErr(err error) {}

type SubjectSearchParameter struct {
	Title         string
	Faculty       string
	AcademicField string
}
