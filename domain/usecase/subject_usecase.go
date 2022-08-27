package usecase

import (
	"github.com/kafugen/ocwcentral/domain/usecase/dto"
)

type SubjectUsecase interface {
	GetById(id string) (*dto.SubjectDTO, error)
	GetByIds(ids []string) ([]*dto.SubjectDTO, error)
}
