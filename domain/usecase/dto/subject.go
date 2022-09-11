package dto

import (
	"time"

	"github.com/kafugen/ocwcentral/model"
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

func NewSubjectDTO(subject *model.Subject) *SubjectDTO {
	videoIds := subject.VideoIds()
	videoIdStrs := make([]string, len(videoIds))
	for i, videoId := range videoIds {
		videoIdStrs[i] = videoId.String()
	}

	resourceIds := subject.ResourceIds()
	resourceIdStrs := make([]string, len(resourceIds))
	for i, resourceId := range resourceIds {
		resourceIdStrs[i] = resourceId.String()
	}

	relatedSubjectIds := subject.RelatedSubjectIds()
	relatedSubjectIdStrs := make([]string, len(relatedSubjectIds))
	for i, relatedSubjectId := range relatedSubjectIds {
		relatedSubjectIdStrs[i] = relatedSubjectId.String()
	}

	var syllabusId string
	if subject.SyllabusId() != nil {
		syllabusId = subject.SyllabusId().String()
	}

	return &SubjectDTO{
		ID:                subject.Id().String(),
		Category:          subject.Category(),
		Title:             subject.Title(),
		VideoIds:          videoIdStrs,
		Location:          subject.Location(),
		ResourceIds:       resourceIdStrs,
		RelatedSubjectIds: relatedSubjectIdStrs,
		Department:        subject.Department(),
		FirstHeldOn:       subject.FirstHeldOn(),
		Faculty:           subject.Faculty(),
		Language:          subject.Language(),
		FreeDescription:   subject.FreeDescription(),
		SyllabusId:        syllabusId,
		Series:            subject.Series(),
		AcademicField:     subject.AcademicField(),
		ThumbnailLink:     subject.ThumbnailLink(),
	}
}
