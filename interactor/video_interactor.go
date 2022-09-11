package interactor

import (
	"fmt"

	"github.com/kafugen/ocwcentral/domain/repository"
	"github.com/kafugen/ocwcentral/domain/usecase/dto"
	"github.com/kafugen/ocwcentral/model"
)

type VideoInteractor struct {
	sR repository.VideoRepository
}

func NewVideoInteractor(sR repository.VideoRepository) *VideoInteractor {
	return &VideoInteractor{sR}
}

func (sI *VideoInteractor) GetByIds(ids []string) ([]*dto.VideoDTO, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	videoIds := make([]model.VideoId, len(ids))
	for i, id := range ids {
		videoId, err := model.NewVideoId(id)
		videoIds[i] = *videoId
		if err != nil {
			return nil, fmt.Errorf("failed on create `VideoId` struct: %w", err)
		}
	}

	videos, err := sI.sR.GetByIds(videoIds)
	if err != nil {
		return nil, fmt.Errorf("failed on executing `GetByIds` of VideoRepository: %w", err)
	}

	videoDTOs := make([]*dto.VideoDTO, len(videos))
	for i, video := range videos {
		videoDTOs[i] = dto.NewVideoDTO(video)
	}
	return videoDTOs, nil
}
