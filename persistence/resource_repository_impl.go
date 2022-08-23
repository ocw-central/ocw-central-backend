package persistence

import (
	"github.com/kafugen/ocwcentral/model"
)

type ResourceRepositoryImpl struct {
}

func NewResourceRepositoryImpl() ResourceRepositoryImpl {
	return ResourceRepositoryImpl{}
}

func (vR ResourceRepositoryImpl) GetByIds(ids []model.ResourceId) ([]*model.Resource, error) {
	panic("not implemented")
}
