package persistence

import (
	"github.com/kafugen/ocwcentral/model"

	"github.com/jmoiron/sqlx"
)

type VideoRepositoryImpl struct {
	db *sqlx.DB
}

func NewVideoRepositoryImpl(db *sqlx.DB) VideoRepositoryImpl {
	return VideoRepositoryImpl{db}
}

func (vR VideoRepositoryImpl) GetByIds(ids []model.VideoId) ([]*model.Video, error) {
	panic("not implemented")
}
