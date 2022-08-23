package model

import (
	"github.com/kafugen/ocwcentral/domain/usecase/dto"
)

type Node interface {
	IsNode()
	GetID() string
}

type Chapter dto.ChapterDTO

func (Chapter) IsNode()            {}
func (this Chapter) GetID() string { return this.ID }

type Resource dto.ResourceDTO

func (Resource) IsNode()            {}
func (this Resource) GetID() string { return this.ID }

type Subject dto.SubjectDTO

func (Subject) IsNode()            {}
func (this Subject) GetID() string { return this.ID }

type Subpage dto.SubpageDTO

func (Subpage) IsNode()            {}
func (this Subpage) GetID() string { return this.ID }

type Syllabus struct {
	dto.SyllabusDTO
	Subpages []Subpage `json:"subpages"`
}

func (Syllabus) IsNode()            {}
func (this Syllabus) GetID() string { return this.ID }

// type Video dto.VideoDTO
type Video struct {
	dto.VideoDTO
	Chapters []Chapter `json:"chapters"`
}

func (Video) IsNode()            {}
func (this Video) GetID() string { return this.ID }

type RelatedSubject dto.SubjectDTO

func (RelatedSubject) IsNode()            {}
func (this RelatedSubject) GetID() string { return this.ID }
