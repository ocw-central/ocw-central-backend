package dto

type ChapterDTO struct {
	ID            string    `json:"id"`
	StartAt       int `json:"startAt"`
	Topic         string    `json:"topic"`
	ThumbnailLink string    `json:"thumbnailLink"`
}
