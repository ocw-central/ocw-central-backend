package dto

import (
	"time"
)

type SubjectDTO struct {
	Id              *[]byte    `db:"id"`
	Category        *string    `db:"category"`
	Title           *string    `db:"title"`
	Location        *string    `db:"location"`
	Department      *string    `db:"department"`
	FirstHeldOn     *time.Time `db:"first_held_on"`
	Faculty         *string    `db:"faculty"`
	Language        *string    `db:"language"`
	FreeDescription *string    `db:"free_description"`
	Series          *string    `db:"series"`
	AcademicField   *string    `db:"academic_field"`
	SyllabusId      *[]byte    `db:"syllabus_id"`
	ThumbnailLink   *string    `db:"thumbnail_link"`
}
