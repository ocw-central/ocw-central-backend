package repository

import (
	"github.com/kafugen/ocwcentral/model"
)

type ResourceRepository interface {
	GetByIds(ids []model.ResourceId) ([]*model.Resource, error)
}
