package persistence

import (
	"github.com/kafugen/ocwcentral/model"

	"github.com/jmoiron/sqlx"
)

type ResourceRepositoryImpl struct {
	db *sqlx.DB
}

func NewResourceRepositoryImpl(db *sqlx.DB) ResourceRepositoryImpl {
	return ResourceRepositoryImpl{db}
}

func (vR ResourceRepositoryImpl) GetByIds(ids []model.ResourceId) ([]*model.Resource, error) {
	panic("not implemented")
}
