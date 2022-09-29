package interactor

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/kafugen/ocwcentral/domain/repository"
	"github.com/kafugen/ocwcentral/domain/usecase/dto"
	"github.com/kafugen/ocwcentral/model"
)

const (
	numRandomSubjects = 12

	// this value is used for GetByRandom() cache
	cacheDefaultExpiration = 2 * time.Hour
	cacheCleanupInterval   = 24 * time.Hour
)

type SubjectInteractor struct {
	sR                 repository.SubjectRepository
	randomSubjectCache *cache.Cache
}

func NewSubjectInteractor(sR repository.SubjectRepository) *SubjectInteractor {
	return &SubjectInteractor{
		sR:                 sR,
		randomSubjectCache: cache.New(cacheDefaultExpiration, cacheCleanupInterval),
	}
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

func (sI SubjectInteractor) GetBySearchParameter(title string, faculty string, academicField string) ([]*dto.SubjectDTO, error) {

	subjects, err := sI.sR.GetBySearchParameter(title, faculty, academicField)
	if err != nil {
		return nil, fmt.Errorf("failed on executing `GetBySearchParameter` of SubjectRepository: %w", err)
	}

	subjectDTOs := make([]*dto.SubjectDTO, len(subjects))
	for i, subject := range subjects {
		subjectDTOs[i] = dto.NewSubjectDTO(subject)
	}
	return subjectDTOs, nil
}

func (sI SubjectInteractor) GetByRandom() ([]*dto.SubjectDTO, error) {
	if cache, found := sI.randomSubjectCache.Get("random-subjects"); found {
		if subjects, ok := cache.([]*dto.SubjectDTO); ok {
			return subjects, nil
		}
	}

	subjects, err := sI.sR.GetByRandom(numRandomSubjects)
	if err != nil {
		return nil, fmt.Errorf("failed on executing `GetByRandom` of SubjectRepository: %w", err)
	}

	subjectDTOs := make([]*dto.SubjectDTO, len(subjects))
	for i, subject := range subjects {
		subjectDTOs[i] = dto.NewSubjectDTO(subject)
	}

	sI.randomSubjectCache.Set("random-subjects", subjectDTOs, cache.DefaultExpiration)
	return subjectDTOs, nil
}
