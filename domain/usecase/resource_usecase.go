package usecase

import (
	"github.com/kafugen/ocwcentral/domain/usecase/dto"
)

type ResourceUsecase interface {
	GetByIds(ids []string) ([]*dto.ResourceDTO, error)
}
