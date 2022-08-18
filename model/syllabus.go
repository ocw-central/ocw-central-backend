package model

import (
	"time"

	"github.com/oklog/ulid"
)

type SyllabusId ulid.ULID

type Syllabus struct {
	id                SyllabusId   `desc:"ID"`
	facultyIds        []FacultyId  `desc:"FacultyIDs"`
	language          string       `desc:"使用言語"`
	subjectNumbering  string       `desc:"科目ナンバリング"`
	academicYear      int16        `desc:"開講年度"`
	semester          string       `desc:"開講期"`
	numCredit         int8         `desc:"単位数"`
	courceFormat      string       `desc:"授業形態"`
	assignedGrade     string       `desc:"配当学年"`
	targetedAudience  string       `desc:"対象学生"`
	dayOfWeek         time.Weekday `desc:"曜日"`
	courcePeriod      string       `desc:"時限"`
	outline           string       `desc:"授業の概要・目的"`
	objective         string       `desc:"到達目標"`
	lessonPlan        string       `desc:"授業計画と内容"`
	gradingMethod     string       `desc:"成績評価の方法・観点"`
	courceRequirement string       `desc:"履修要件"`
	outClassLearning  string       `desc:"授業外学習（予習・復習）等"`
	reference         string       `desc:"教科書・参考書等"`
	remark            string       `desc:"備考"`
	subpageIds        []SubpageId  `desc:"サブページIDs"`
}
