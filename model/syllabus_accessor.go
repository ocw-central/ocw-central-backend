package model

func (s *Syllabus) Id() SyllabusId {
	return s.id
}
func (s *Syllabus) Faculty() string {
	return s.faculty
}
func (s *Syllabus) Language() string {
	return s.language
}
func (s *Syllabus) SubjectNumbering() string {
	return s.subjectNumbering
}
func (s *Syllabus) AcademicYear() int16 {
	return s.academicYear
}
func (s *Syllabus) Semester() string {
	return s.semester
}
func (s *Syllabus) NumCredit() int8 {
	return s.numCredit
}
func (s *Syllabus) CourceFormat() string {
	return s.courceFormat
}
func (s *Syllabus) AssignedGrade() string {
	return s.assignedGrade
}
func (s *Syllabus) TargetedAudience() string {
	return s.targetedAudience
}
func (s *Syllabus) CourceDayPeriod() string {
	return s.courceDayPeriod
}
func (s *Syllabus) Outline() string {
	return s.outline
}
func (s *Syllabus) Objective() string {
	return s.objective
}
func (s *Syllabus) LessonPlan() string {
	return s.lessonPlan
}
func (s *Syllabus) GradingMethod() string {
	return s.gradingMethod
}
func (s *Syllabus) CourceRequirement() string {
	return s.courceRequirement
}
func (s *Syllabus) OutClassLearning() string {
	return s.outClassLearning
}
func (s *Syllabus) Reference() string {
	return s.reference
}
func (s *Syllabus) Remark() string {
	return s.remark
}
func (s *Syllabus) SubpageIds() []Subpage {
	return s.subpages
}
