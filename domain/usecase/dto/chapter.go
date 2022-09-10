package dto

import (
	"github.com/kafugen/ocwcentral/model"
)

type ChapterDTO struct {
	ID            string `json:"id"`
	StartAt       int    `json:"startAt"`
	Topic         string `json:"topic"`
	ThumbnailLink string `json:"thumbnailLink"`
}

func NewChapterDTO(chapter *model.Chapter) *ChapterDTO {
	return &ChapterDTO{
		ID:            chapter.Id().String(),
		StartAt:       chapter.StartAt(),
		Topic:         chapter.Topic(),
		ThumbnailLink: chapter.ThumbnailLink(),
	}
}
