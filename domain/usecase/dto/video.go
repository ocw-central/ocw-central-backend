package dto

import (
	"time"

	"github.com/kafugen/ocwcentral/model"
)

type VideoDTO struct {
	ID          string       `json:"id"`
	Title       string       `json:"title"`
	Ordering    int          `json:"ordering"`
	Link        string       `json:"link"`
	Chapters    []ChapterDTO `json:"chapters"`
	Faculty     string       `json:"faculty"`
	LecturedOn  time.Time    `json:"lecturedOn"`
	VideoLength int          `json:"videoLength"`
	Language    string       `json:"language"`
}

func NewVideoDTO(video *model.Video) *VideoDTO {
	chapterDTOs := make([]ChapterDTO, len(video.Chapters()))
	for j, chapter := range video.Chapters() {
		chapterDTOs[j] = *NewChapterDTO(&chapter)
	}

	return &VideoDTO{
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
