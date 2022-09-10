package interactor

import (
	"fmt"

	"github.com/kafugen/ocwcentral/domain/repository"
	"github.com/kafugen/ocwcentral/domain/usecase/dto"
	"github.com/kafugen/ocwcentral/model"
)

type SubjectInteractor struct {
	sR repository.SubjectRepository
}

func NewSubjectInteractor(sR repository.SubjectRepository) *SubjectInteractor {
	return &SubjectInteractor{sR}
}

func (sI *SubjectInteractor) GetById(id string) (*dto.SubjectDTO, error) {
	subjectId, err := model.NewSubjectId(id)
	if err != nil {
		return nil, fmt.Errorf("failed on create `SubjectId` struct: %w", err)
	}

	subject, err := sI.sR.GetById(*subjectId)
	if err != nil {
		return nil, fmt.Errorf("failed on executing `GetById` of SubjectRepository: %w", err)
	}

	return dto.NewSubjectDTO(subject), nil
}

func (sI *SubjectInteractor) GetByIds(ids []string) ([]*dto.SubjectDTO, error) {
	subjectIds := make([]model.SubjectId, len(ids))
	for i, id := range ids {
		subjectId, err := model.NewSubjectId(id)
		if err != nil {
			return nil, fmt.Errorf("failed on create `SubjectId`: %w", err)
		}
		subjectIds[i] = *subjectId
	}

	subjects, err := sI.sR.GetByIds(subjectIds)
	if err != nil {
		return nil, fmt.Errorf("failed on executing `GetByIds` of SubjectRepository: %w", err)
	}

	subjectDTOs := make([]*dto.SubjectDTO, len(subjects))
	for i, subject := range subjects {
		subjectDTOs[i] = dto.NewSubjectDTO(subject)
	}
	return subjectDTOs, nil
}
