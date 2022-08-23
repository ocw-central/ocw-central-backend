package persistence

import (
	"github.com/kafugen/ocwcentral/model"
)

type SyllabusRepositoryImpl struct {
}

func NewSyllabusRepositoryImpl() SyllabusRepositoryImpl {
	return SyllabusRepositoryImpl{}
}

func (sR SyllabusRepositoryImpl) GetById(id model.SyllabusId) (*model.Syllabus, error) {
	panic("not implemented")
}
