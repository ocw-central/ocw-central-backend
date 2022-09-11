package usecase

type AcademicFieldUsecase interface {
	Get() ([]string, error)
}
