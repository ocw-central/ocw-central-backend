package model

import (
	"time"
)

func (s *Subject) Id() SubjectId {
	return s.id
}
func (s *Subject) Category() string {
	return s.category
}
func (s *Subject) Title() string {
	return s.title
}
func (s *Subject) VideoIds() []VideoId {
	return s.videoIds
}
func (s *Subject) Location() string {
	return s.location
}
func (s *Subject) ResourceIds() []ResourceId {
	return s.resourceIds
}
func (s *Subject) RelatedSubjectIds() []SubjectId {
	return s.relatedSubjectIds
}
func (s *Subject) Department() string {
	return s.department
}
func (s *Subject) FirstHeldOn() time.Time {
	return s.firstHeldOn
}
func (s *Subject) Faculty() string {
	return s.faculty
}
func (s *Subject) Language() string {
	return s.language
}
func (s *Subject) FreeDescription() string {
	return s.freeDescription
}
func (s *Subject) SyllabusId() *SyllabusId {
	return s.syllabusId
}
func (s *Subject) Series() string {
	return s.series
}
func (s *Subject) AcademicField() string {
	return s.academicField
}
func (s *Subject) ThumbnailLink() string {
	return s.thumbnailLink
}
