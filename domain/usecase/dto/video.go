package dto

import (
	"time"

	"github.com/kafugen/ocwcentral/model"
)

type VideoDTO struct {
	ID            string           `json:"id"`
	Title         string           `json:"title"`
	Ordering      int              `json:"ordering"`
	Link          string           `json:"link"`
	Chapters      []ChapterDTO     `json:"chapters"`
	Faculty       string           `json:"faculty"`
	LecturedOn    time.Time        `json:"lecturedOn"`
	VideoLength   int              `json:"videoLength"`
	Language      string           `json:"language"`
	Transcription string           `json:"transcription"`
	Translations  []TranslationDTO `json:"translations"`
}

func NewVideoDTO(video *model.Video) *VideoDTO {
	chapterDTOs := make([]ChapterDTO, len(video.Chapters()))
	for j, chapter := range video.Chapters() {
		chapterDTOs[j] = *NewChapterDTO(&chapter)
	}

	translationDTOs := make([]TranslationDTO, len(video.Translations()))
	for j, translation := range video.Translations() {
		translationDTOs[j] = *NewTranslationDTO(&translation)
	}

	return &VideoDTO{
		ID:            video.Id().String(),
		Title:         video.Title(),
		Ordering:      video.Ordering(),
		Link:          video.Link(),
		Chapters:      chapterDTOs,
		Faculty:       video.Faculty(),
		LecturedOn:    video.LecturedOn(),
		VideoLength:   int(video.VideoLength().Seconds()),
		Language:      video.Language(),
		Transcription: video.Transcription(),
		Translations:  translationDTOs,
	}
}
