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
	// this value is used for GetByRandom() cache
	cacheDefaultExpiration = 2 * time.Hour
	cacheCleanupInterval   = 24 * time.Hour
)

type SubjectInteractor struct {
	sR                 repository.SubjectRepository
	vR                 repository.VideoRepository
	randomSubjectCache *cache.Cache
}

func NewSubjectInteractor(sR repository.SubjectRepository, vR repository.VideoRepository) *SubjectInteractor {
	return &SubjectInteractor{
		sR:                 sR,
		vR:                 vR,
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

func (sI SubjectInteractor) GetByRandom(category string, series string, academicField string, numRandomSubjects int) ([]*dto.SubjectDTO, error) {
	if cache, found := sI.randomSubjectCache.Get("random-subjects"); found {
		if subjects, ok := cache.([]*dto.SubjectDTO); ok {
			return subjects, nil
		}
	}

	subjects, err := sI.sR.GetByRandom(category, series, academicField, numRandomSubjects)
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

func (sI SubjectInteractor) GetByVideoSearchParameter(title string, faculty string) ([]*dto.SubjectWithSpecifiedVideosDTO, error) {
	videos, err := sI.vR.GetBySearchParameter(title, faculty)
	if err != nil {
		return nil, fmt.Errorf("failed on executing `GetBySearchParameter` of VideoRepository: %w", err)
	}

	m := map[string]*model.Video{}
	videoIds := make([]model.VideoId, len(videos))
	for _, video := range videos {
		m[video.Id().String()] = video
		videoIds = append(videoIds, video.Id())
	}

	subjects, err := sI.sR.GetByVideoIds(videoIds)
	if err != nil {
		return nil, fmt.Errorf("failed on executing `GetByVideoIds` of SubjectRepository: %w", err)
	}

	subjectWithSpecifiedVideosDTOs := make([]*dto.SubjectWithSpecifiedVideosDTO, len(subjects))
	for i, subject := range subjects {
		videos := make([]*model.Video, 0, 20)
		for _, videoId := range subject.VideoIds() {
			_, ok := m[videoId.String()]
			if ok {
				videos = append(videos, m[videoId.String()])
			}
		}
		subjectWithSpecifiedVideosDTOs[i] = dto.NewSubjectWithSpecifiedVideosDTO(subject, videos)
	}

	return subjectWithSpecifiedVideosDTOs, nil
}
