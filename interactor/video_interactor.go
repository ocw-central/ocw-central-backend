package interactor

import (
	"github.com/kafugen/ocwcentral/domain/repository"
	"github.com/kafugen/ocwcentral/domain/usecase/dto"
)

type VideoInteractor struct {
	sR repository.VideoRepository
}

func NewVideoInteractor(sR repository.VideoRepository) VideoInteractor {
	return VideoInteractor{sR}
}

func (sI VideoInteractor) GetByIds(ids []string) ([]*dto.VideoDTO, error) {
	panic("not implemented")
}
