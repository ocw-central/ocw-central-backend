package dto

type SyllabusSubpageDTO struct {
	Id                *[]byte `db:"id"`
	Faculty           *string `db:"faculty"`
	Language          *string `db:"language"`
	SubjectNumbering  *string `db:"subject_numbering"`
	AcademicYear      *string `db:"academic_year"`
	Semester          *string `db:"semester"`
	NumCredit         *int8   `db:"num_credit"`
	CourseFormat      *string `db:"course_format"`
	AssignedGrade     *string `db:"assigned_grade"`
	TargettedAudience *string `db:"targetted_audience"`
	CourseDayPeriod   *string `db:"course_day_period"`
	Outline           *string `db:"outline"`
	Objective         *string `db:"objective"`
	LessonPlan        *string `db:"lesson_plan"`
	GradingMethod     *string `db:"grading_method"`
	CourseRequirement *string `db:"course_requirement"`
	OutclassLearning  *string `db:"outclass_learning"`
	Reference         *string `db:"reference"`
	Remark            *string `db:"remark"`
	SubpageId         *[]byte `db:"subpage_ids"`
	Content           *string `db:"content"`
}
