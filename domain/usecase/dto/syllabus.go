package dto

import (
	"github.com/kafugen/ocwcentral/model"
)

type SyllabusDTO struct {
	ID                string       `json:"id"`
	Faculty           string       `json:"faculty"`
	Language          string       `json:"language"`
	SubjectNumbering  string       `json:"subjectNumbering"`
	AcademicYear      int16        `json:"academicYear"`
	Semester          string       `json:"semester"`
	NumCredit         int8         `json:"numCredit"`
	CourseFormat      string       `json:"courceFormat"`
	AssignedGrade     string       `json:"assignedGrade"`
	TargetedAudience  string       `json:"targetedAudience"`
	CourseDayPeriod   string       `json:"courceDayPeriod"`
	Outline           string       `json:"outline"`
	Objective         string       `json:"objective"`
	LessonPlan        string       `json:"lessonPlan"`
	GradingMethod     string       `json:"gradingMethod"`
	CourseRequirement string       `json:"courceRequirement"`
	OutClassLearning  string       `json:"outClassLearning"`
	Reference         string       `json:"reference"`
	Remark            string       `json:"remark"`
	Subpages          []SubpageDTO `json:"subpages"`
}

func NewSyllabusDTO(syllabus *model.Syllabus) *SyllabusDTO {
	subpageDTOs := make([]SubpageDTO, len(syllabus.Subpages()))
	for i, subpage := range syllabus.Subpages() {
		subpageDTOs[i] = *NewSubpageDTO(&subpage)
	}

	return &SyllabusDTO{
		ID:                syllabus.Id().String(),
		Faculty:           syllabus.Faculty(),
		Language:          syllabus.Language(),
		SubjectNumbering:  syllabus.SubjectNumbering(),
		AcademicYear:      syllabus.AcademicYear(),
		Semester:          syllabus.Semester(),
		NumCredit:         syllabus.NumCredit(),
		CourseFormat:      syllabus.CourseFormat(),
		AssignedGrade:     syllabus.AssignedGrade(),
		TargetedAudience:  syllabus.TargetedAudience(),
		CourseDayPeriod:   syllabus.CourseDayPeriod(),
		Outline:           syllabus.Outline(),
		Objective:         syllabus.Objective(),
		LessonPlan:        syllabus.LessonPlan(),
		GradingMethod:     syllabus.GradingMethod(),
		CourseRequirement: syllabus.CourseRequirement(),
		OutClassLearning:  syllabus.OutClassLearning(),
		Reference:         syllabus.Reference(),
		Remark:            syllabus.Remark(),
		Subpages:          subpageDTOs,
	}
}
