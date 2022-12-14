package model

import (
	"github.com/oklog/ulid"
)

type SyllabusId ulid.ULID

func NewSyllabusId[T string | []byte](ulidExp T) (*SyllabusId, error) {
	var scannedULID ulid.ULID
	err := scannedULID.Scan(ulidExp)
	if err != nil {
		return nil, err
	}
	syllabusId := SyllabusId(scannedULID)
	return &syllabusId, nil
}

func (s SyllabusId) String() string {
	return ulid.ULID(s).String()
}

func (s SyllabusId) ByteSlice() []byte {
	bytes := [16]byte(s)
	return bytes[:]
}

type Syllabus struct {
	id                SyllabusId `desc:"ID"`
	faculty           string     `desc:"教員の氏名と所属職位"`
	language          string     `desc:"使用言語"`
	subjectNumbering  string     `desc:"科目ナンバリング"`
	academicYear      int16      `desc:"開講年度"`
	semester          string     `desc:"開講期"`
	numCredit         int8       `desc:"単位数"`
	courseFormat      string     `desc:"授業形態"`
	assignedGrade     string     `desc:"配当学年"`
	targetedAudience  string     `desc:"対象学生"`
	courseDayPeriod   string     `desc:"曜時限"`
	outline           string     `desc:"授業の概要・目的"`
	objective         string     `desc:"到達目標"`
	lessonPlan        string     `desc:"授業計画と内容"`
	gradingMethod     string     `desc:"成績評価の方法・観点"`
	courseRequirement string     `desc:"履修要件"`
	outClassLearning  string     `desc:"授業外学習（予習・復習）等"`
	reference         string     `desc:"教科書・参考書等"`
	remark            string     `desc:"備考"`
	subpages          []Subpage  `desc:"サブページs"`
}

func NewSyllabusFromRepository(
	id SyllabusId,
	faculty string,
	language string,
	subjectNumbering string,
	academicYear int16,
	semester string,
	numCredit int8,
	courseFormat string,
	assignedGrade string,
	targetedAudience string,
	courseDayPeriod string,
	outline string,
	objective string,
	lessonPlan string,
	gradingMethod string,
	courseRequirement string,
	outClassLearning string,
	reference string,
	remark string,
	subpages []Subpage,
) *Syllabus {
	return &Syllabus{
		id:                id,
		faculty:           faculty,
		language:          language,
		subjectNumbering:  subjectNumbering,
		academicYear:      academicYear,
		semester:          semester,
		numCredit:         numCredit,
		courseFormat:      courseFormat,
		assignedGrade:     assignedGrade,
		targetedAudience:  targetedAudience,
		courseDayPeriod:   courseDayPeriod,
		outline:           outline,
		objective:         objective,
		lessonPlan:        lessonPlan,
		gradingMethod:     gradingMethod,
		courseRequirement: courseRequirement,
		outClassLearning:  outClassLearning,
		reference:         reference,
		remark:            remark,
		subpages:          subpages,
	}
}
