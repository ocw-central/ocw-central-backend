package interactor

import (
	"fmt"

	"github.com/kafugen/ocwcentral/domain/repository"
)

type AcademicFieldInteractor struct {
	afR repository.AcademicFieldRepository
}

func NewAcademicFieldInteractor(afR repository.AcademicFieldRepository) *AcademicFieldInteractor {
	return &AcademicFieldInteractor{afR: afR}
}

func (afI *AcademicFieldInteractor) Get() ([]string, error) {
	academicFields, err := afI.afR.Get()
	if err != nil {
		return nil, fmt.Errorf("failed on executing `Get` of AcademicFieldRepository: %w", err)
	}
	return academicFields, nil
}
