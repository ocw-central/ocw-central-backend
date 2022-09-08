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

	return convertModelToDTO(subject), nil
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
		subjectDTOs[i] = convertModelToDTO(subject)
	}
	return subjectDTOs, nil
}

func convertModelToDTO(subject *model.Subject) *dto.SubjectDTO {
	videoIds := subject.VideoIds()
	videoIdStrs := make([]string, len(videoIds))
	for i, videoId := range videoIds {
		videoIdStrs[i] = videoId.String()
	}

	resourceIds := subject.ResourceIds()
	resourceIdStrs := make([]string, len(resourceIds))
	for i, resourceId := range resourceIds {
		resourceIdStrs[i] = resourceId.String()
	}

	relatedSubjectIds := subject.RelatedSubjectIds()
	relatedSubjectIdStrs := make([]string, len(relatedSubjectIds))
	for i, relatedSubjectId := range relatedSubjectIds {
		relatedSubjectIdStrs[i] = relatedSubjectId.String()
	}

	var syllabusId string
	if subject.SyllabusId() != nil {
		syllabusId = subject.SyllabusId().String()
	}

	return &dto.SubjectDTO{
		ID:                subject.Id().String(),
		Category:          subject.Category(),
		Title:             subject.Title(),
		VideoIds:          videoIdStrs,
		Location:          subject.Location(),
		ResourceIds:       resourceIdStrs,
		RelatedSubjectIds: relatedSubjectIdStrs,
		Department:        subject.Department(),
		FirstHeldOn:       subject.FirstHeldOn(),
		Faculty:           subject.Faculty(),
		Language:          subject.Language(),
		FreeDescription:   subject.FreeDescription(),
		SyllabusId:        syllabusId,
		Series:            subject.Series(),
		AcademicField:     subject.AcademicField(),
	}
}
