package model

import (
	"github.com/oklog/ulid"
)

type SubjectId struct {
	id ulid.ULID `desc:"ID"`
}

type Subject struct {
	id                SubjectId    `desc:"科目.ID"`
	category          string       `desc:"科目.カテゴリ"`
	title             string       `desc:"科目.名前"`
	videoIds          []*VideoId   `desc:"科目.VideoIds"`
	location          string       `desc:"科目.開催場所"`
	pdfIds            []*PdfId     `desc:"科目.PdfIds"`
	relatedSubjectIds []*SubjectId `desc:"科目.関連科目IDs"`
	department        string       `desc:"科目.開講部局名"`
	firstHeldOn       string       `desc:"科目.開催日"`
	facultyIds        []*FacultyId `desc:"科目(シラバス).FacultyIds"`
	language          string       `desc:"科目(シラバス).使用言語"`
	subjectNumbering  string       `desc:"シラバス.科目ナンバリング"`
	academicYear      string       `desc:"シラバス.開講年度"`
	semester          string       `desc:"シラバス.開講期"`
	numCredit         int8         `desc:"シラバス.単位数"`
	courceFormat      string       `desc:"シラバス.授業形態"`
	assignedGrade     string       `desc:"シラバス.配当学年"`
	targetedAudience  string       `desc:"シラバス.対象学生"`
	dayOfWeek         string       `desc:"シラバス.曜日"`
	courcePeriod      string       `desc:"シラバス.時限"`
	outline           string       `desc:"シラバス.授業の概要・目的"`
	objective         string       `desc:"シラバス.到達目標"`
	lessonPlan        string       `desc:"シラバス.授業計画と内容"`
	gradingMethod     string       `desc:"シラバス.成績評価の方法・観点"`
	courceRequirement string       `desc:"シラバス.履修要件"`
	outClassLearning  string       `desc:"シラバス.授業外学習（予習・復習）等"`
	reference         string       `desc:"シラバス.教科書・参考書等"`
	remark            string       `desc:"シラバス.備考"`
	subpageIds        []*Subpage   `desc:"科目.サブページIDs"`
}
