package usecase

import (
	"github.com/kafugen/ocwcentral/domain/usecase/dto"
)

type SubjectUsecase interface {
	GetById(id string) (*dto.SubjectDTO, error)
}
