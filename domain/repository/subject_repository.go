package repository

import (
	"github.com/kafugen/ocwcentral/model"
)

type SubjectRepository interface {
	GetById(id model.SubjectId) (*model.Subject, error)
	GetByIds(ids []model.SubjectId) ([]*model.Subject, error)
}
