package interactor

import (
	"github.com/kafugen/ocwcentral/domain/repository"
	"github.com/kafugen/ocwcentral/domain/usecase/dto"
)

type ResourceInteractor struct {
	sR repository.ResourceRepository
}

func NewResourceInteractor(sR repository.ResourceRepository) ResourceInteractor {
	return ResourceInteractor{sR}
}

func (sI ResourceInteractor) GetByIds(ids []string) ([]*dto.ResourceDTO, error) {
	panic("not implemented")
}
