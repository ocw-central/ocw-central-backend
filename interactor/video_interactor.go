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
		chapterDTOs := make([]dto.ChapterDTO, len(video.Chapters()))
		for j, chapter := range video.Chapters() {
			chapterDTOs[j] = dto.ChapterDTO{
				ID:            chapter.Id().String(),
				StartAt:       chapter.StartAt(),
				Topic:         chapter.Topic(),
				ThumbnailLink: chapter.ThumbnailLink(),
			}
		}

		videoDTOs[i] = &dto.VideoDTO{
			ID:          video.Id().String(),
			Title:       video.Title(),
			Ordering:    video.Ordering(),
			Link:        video.Link(),
			Chapters:    chapterDTOs,
			Faculty:     video.Faculty(),
			LecturedOn:  video.LecturedOn(),
			VideoLength: int(video.VideoLength().Seconds()),
			Language:    video.Language(),
		}
	}
	return videoDTOs, nil
}
