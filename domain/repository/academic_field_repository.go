package repository

type AcademicFieldRepository interface {
	Get() ([]string, error)
}
