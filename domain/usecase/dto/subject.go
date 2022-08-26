package dto

import (
	"time"
)

type SubjectDTO struct {
	ID                string
	Category          string
	Title             string
	VideoIds          []string
	Location          string
	ResourceIds       []string
	RelatedSubjectIds []string
	Department        string
	FirstHeldOn       time.Time
	Faculty           string
	Language          string
	FreeDescription   string
	SyllabusId        string
	Series            string
	AcademicField     string
	ThumbnailLink     string
}
