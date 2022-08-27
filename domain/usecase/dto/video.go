package dto

import (
	"time"
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
