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

func NewSubjectInteractor(sR repository.SubjectRepository) SubjectInteractor {
	return SubjectInteractor{sR}
}

func (sI SubjectInteractor) GetById(id string) (*dto.SubjectDTO, error) {
	subjectId, err := model.NewSubjectId(id)
	if err != nil {
		return nil, fmt.Errorf("failed on create `SubjectId` struct: %w", err)
	}

	subject, err := sI.sR.GetById(*subjectId)
	if err != nil {
		return nil, fmt.Errorf("failed on executing `GetById` of SubjectRepository: %w", err)
	}

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

	subjectDTO := dto.SubjectDTO{
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
	return &subjectDTO, nil
}
