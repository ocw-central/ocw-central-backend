package repository

import (
	"github.com/kafugen/ocwcentral/model"
)

type SubjectRepository interface {
	GetById(id model.SubjectId) (*model.Subject, error)
	GetByIds(ids []model.SubjectId) ([]*model.Subject, error)
	GetBySearchParameter(title string, faculty string, academicField string) ([]*model.Subject, error)
	GetByRandom(series string, academicField string, numSubjects int) ([]*model.Subject, error)
	GetByVideoIds(videoIds []model.VideoId) ([]*model.Subject, error)
}
