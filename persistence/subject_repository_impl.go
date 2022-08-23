package persistence

import (
	"github.com/kafugen/ocwcentral/model"
)

type SubjectRepositoryImpl struct {
}

func NewSubjectRepositoryImpl() SubjectRepositoryImpl {
	return SubjectRepositoryImpl{}
}

func (sR SubjectRepositoryImpl) GetById(id model.SubjectId) (*model.Subject, error) {
	panic("not implemented")
}

func (sR SubjectRepositoryImpl) GetByIds(id []model.SubjectId) ([]*model.Subject, error) {
	panic("not implemented")
}
