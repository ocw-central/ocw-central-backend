package repository

import (
	"github.com/kafugen/ocwcentral/model"
)

type ResourceRepository interface {
	GetById(id model.ResourceId) (*model.Resource, error)
	GetByIds(ids []model.ResourceId) ([]*model.Resource, error)
}
