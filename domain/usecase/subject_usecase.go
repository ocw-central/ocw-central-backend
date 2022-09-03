package usecase

import (
	"github.com/kafugen/ocwcentral/domain/usecase/dto"
	"github.com/kafugen/ocwcentral/utils"
)

type SubjectUsecase interface {
	GetById(id string) (*dto.SubjectDTO, error)
	GetByIds(ids []string) ([]*dto.SubjectDTO, error)
	GetBySearchParameter(searchParameter utils.SubjectSearchParameter) ([]*dto.SubjectDTO, error)
}
