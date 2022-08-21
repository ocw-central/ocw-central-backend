package model

import (
	"time"

	"github.com/oklog/ulid"
)

type SubjectId ulid.ULID

type Subject struct {
	id                SubjectId    `desc:"ID"`
	category          string       `desc:"カテゴリ"`
	title             string       `desc:"名前"`
	videoIds          []VideoId    `desc:"VideoIds"`
	location          string       `desc:"開催場所"`
	resourceIds       []ResourceId `desc:"ResourceIds"`
	relatedSubjectIds []SubjectId  `desc:"関連科目IDs"`
	department        string       `desc:"開講部局名"`
	firstHeldOn       time.Time    `desc:"開催日"`
	faculty.          string       `desc:"FacultyIds"`
	language          string       `desc:"使用言語"`
	freeDescription   string       `desc:"自由な説明"`
	syllabusId        SyllabusId   `desc:"SyllabusId"`
	series            string       `desc:"シリーズ"`
	academicField     string       `desc:"分野名"`
}
