package persistence

import (
	"github.com/kafugen/ocwcentral/model"

	"github.com/jmoiron/sqlx"
)

type SubjectRepositoryImpl struct {
	db *sqlx.DB
}

func NewSubjectRepositoryImpl(db *sqlx.DB) SubjectRepositoryImpl {
	return SubjectRepositoryImpl{db}
}

func (sR SubjectRepositoryImpl) GetById(id model.SubjectId) (*model.Subject, error) {
	panic("not implemented")
}

func (sR SubjectRepositoryImpl) GetByIds(id []model.SubjectId) ([]*model.Subject, error) {
	panic("not implemented")
}
