package usecase

import (
	"github.com/kafugen/ocwcentral/domain/usecase/dto"
)

type VideoUsecase interface {
	GetByIds(ids []string) ([]*dto.VideoDTO, error)
}
