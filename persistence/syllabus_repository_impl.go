package persistence

import (
	"github.com/kafugen/ocwcentral/model"

	"github.com/jmoiron/sqlx"
)

type SyllabusRepositoryImpl struct {
	db *sqlx.DB
}

func NewSyllabusRepositoryImpl(db *sqlx.DB) SyllabusRepositoryImpl {
	return SyllabusRepositoryImpl{db}
}

func (sR SyllabusRepositoryImpl) GetById(id model.SyllabusId) (*model.Syllabus, error) {
	panic("not implemented")
}
