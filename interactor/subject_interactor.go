package interactor

import (
	"github.com/kafugen/ocwcentral/domain/repository"
	"github.com/kafugen/ocwcentral/domain/usecase/dto"
)

type SubjectInteractor struct {
	sR repository.SubjectRepository
}

func NewSubjectInteractor(sR repository.SubjectRepository) SubjectInteractor {
	return SubjectInteractor{sR}
}

func (sI SubjectInteractor) GetById(id string) (*dto.SubjectDTO, error) {
	panic("not implemented")
}
