package dto

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
