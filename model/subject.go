package model

import (
	"time"

	"github.com/oklog/ulid"
)

type SubjectId ulid.ULID

// NewSubjectId generate a new SubjectId from a string or []byte representation of a ULID.
func NewSubjectId[T string | []byte](ulidExp T) (*SubjectId, error) {
	var scannedULID ulid.ULID
	err := scannedULID.Scan(ulidExp)
	if err != nil {
		return nil, err
	}
	subjectId := SubjectId(scannedULID)
	return &subjectId, nil
}

func (s SubjectId) String() string {
	return ulid.ULID(s).String()
}

func (s SubjectId) ByteSlice() []byte {
	bytes := [16]byte(s)
	return bytes[:]
}

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
	faculty           string       `desc:"教員の氏名と所属職位"`
	language          string       `desc:"使用言語"`
	freeDescription   string       `desc:"自由な説明"`
	syllabusId        *SyllabusId  `desc:"SyllabusId"`
	series            string       `desc:"シリーズ"`
	academicField     string       `desc:"分野名"`
	thumbnailLink     string       `desc:"サムネイルのリンク"`
}

func NewSubjectFromRepository(
	id SubjectId,
	category string,
	title string,
	videoIds []VideoId,
	location string,
	resourceIds []ResourceId,
	relatedSubjectIds []SubjectId,
	department string,
	firstHeldOn time.Time,
	faculty string,
	language string,
	freeDescription string,
	syllabusId *SyllabusId,
	series string,
	academicField string,
	thumbnailLink string,
) *Subject {
	return &Subject{id, category, title, videoIds, location, resourceIds,
		relatedSubjectIds, department, firstHeldOn, faculty, language,
		freeDescription, syllabusId, series, academicField, thumbnailLink,
	}
}
