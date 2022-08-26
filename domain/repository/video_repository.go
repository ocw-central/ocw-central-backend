package repository

import (
	"github.com/kafugen/ocwcentral/model"
)

type VideoRepository interface {
	GetByIds(ids []model.VideoId) ([]*model.Video, error)
}
