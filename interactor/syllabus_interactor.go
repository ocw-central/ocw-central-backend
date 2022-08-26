package interactor

import (
	"github.com/kafugen/ocwcentral/domain/repository"
	"github.com/kafugen/ocwcentral/domain/usecase/dto"
)

type SyllabusInteractor struct {
	sR repository.SyllabusRepository
}

func NewSyllabusInteractor(sR repository.SyllabusRepository) SyllabusInteractor {
	return SyllabusInteractor{sR}
}

func (sI SyllabusInteractor) GetById(id []string) (*dto.SyllabusDTO, error) {
	panic("not implemented")
}
