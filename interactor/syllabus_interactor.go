package interactor

import (
	"fmt"

	"github.com/kafugen/ocwcentral/domain/repository"
	"github.com/kafugen/ocwcentral/domain/usecase/dto"
	"github.com/kafugen/ocwcentral/model"
)

type SyllabusInteractor struct {
	sR repository.SyllabusRepository
}

func NewSyllabusInteractor(sR repository.SyllabusRepository) *SyllabusInteractor {
	return &SyllabusInteractor{sR}
}

func (sI *SyllabusInteractor) GetById(id string) (*dto.SyllabusDTO, error) {
	syllabusId, err := model.NewSyllabusId(id)
	if err != nil {
		return nil, fmt.Errorf("failed to create `SyllabusId` struct: %w", err)
	}

	syllabus, err := sI.sR.GetById(*syllabusId)
	if err != nil {
		return nil, fmt.Errorf("failed on executing `GetById` of SyllabusRepository: %w", err)
	}

	return dto.NewSyllabusDTO(syllabus), nil
}
