package dto

import "time"

type VideoChapterDTO struct {
	Id            *[]byte    `db:"id"`
	SubjectId     *string    `db:"subject_id"`
	Title         *string    `db:"title"`
	Faculty       *string    `db:"faculty"`
	Ordering      *int       `db:"ordering"`
	Link          *string    `db:"link"`
	LecturedOn    *time.Time `db:"lectured_on"`
	VideoLength   *int       `db:"video_length"`
	Language      *string    `db:"language"`
	ChapterId     *[]byte    `db:"chapter_id"`
	StartAt       *int       `db:"start_at"`
	Topic         *string    `db:"topic"`
	ThumbnailLink *string    `db:"thumbnail_link"`
	Transcription *string    `db:"transcription"`
}
