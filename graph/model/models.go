package model

import (
	"time"

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
	ID                string    `json:"id"`
	Faculty           string    `json:"faculty"`
	Language          string    `json:"language"`
	SubjectNumbering  string    `json:"subjectNumbering"`
	AcademicYear      int16     `json:"academicYear"`
	Semester          string    `json:"semester"`
	NumCredit         int8      `json:"numCredit"`
	CourseFormat      string    `json:"courseFormat"`
	AssignedGrade     string    `json:"assignedGrade"`
	TargetedAudience  string    `json:"targetedAudience"`
	CourseDayPeriod   string    `json:"courseDayPeriod"`
	Outline           string    `json:"outline"`
	Objective         string    `json:"objective"`
	LessonPlan        string    `json:"lessonPlan"`
	GradingMethod     string    `json:"gradingMethod"`
	CourseRequirement string    `json:"courseRequirement"`
	OutClassLearning  string    `json:"outClassLearning"`
	Reference         string    `json:"reference"`
	Remark            string    `json:"remark"`
	Subpages          []Subpage `json:"subpages"`
}

func (Syllabus) IsNode()            {}
func (this Syllabus) GetID() string { return this.ID }

// type Video dto.VideoDTO
type Video struct {
	ID            string    `json:"id"`
	Ordering      int       `json:"ordering"`
	Title         string    `json:"title"`
	Link          string    `json:"link"`
	Chapters      []Chapter `json:"chapters"`
	Faculty       string    `json:"faculty"`
	LecturedOn    time.Time `json:"lecturedOn"`
	VideoLength   int       `json:"videoLength"`
	Language      string    `json:"language"`
	Transcription string    `json:"transcription"`
}

func (Video) IsNode()            {}
func (this Video) GetID() string { return this.ID }

func NewVideo(videoDTO *dto.VideoDTO) *Video {
	chapters := make([]Chapter, len(videoDTO.Chapters))
	for i, chapter := range videoDTO.Chapters {
		chapters[i] = Chapter(chapter)
	}
	return &Video{
		ID:          videoDTO.ID,
		Ordering:    videoDTO.Ordering,
		Title:       videoDTO.Title,
		Link:        videoDTO.Link,
		Chapters:    chapters,
		Faculty:     videoDTO.Faculty,
		LecturedOn:  videoDTO.LecturedOn,
		VideoLength: videoDTO.VideoLength,
		Language:    videoDTO.Language,
	}
}

type RelatedSubject dto.SubjectDTO

func (RelatedSubject) IsNode()            {}
func (this RelatedSubject) GetID() string { return this.ID }
