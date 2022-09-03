package repository

import (
	"github.com/kafugen/ocwcentral/model"
	"github.com/kafugen/ocwcentral/utils"
)

type SubjectRepository interface {
	GetById(id model.SubjectId) (*model.Subject, error)
	GetByIds(ids []model.SubjectId) ([]*model.Subject, error)
	GetBySearchParameter(searchParameter utils.SubjectSearchParameter) ([]*model.Subject, error)
}
