package persistence

import (
	"github.com/kafugen/ocwcentral/model"
)

type VideoRepositoryImpl struct {
}

func NewVideoRepositoryImpl() VideoRepositoryImpl {
	return VideoRepositoryImpl{}
}

func (vR VideoRepositoryImpl) GetByIds(ids []model.VideoId) ([]*model.Video, error) {
	panic("not implemented")
}
