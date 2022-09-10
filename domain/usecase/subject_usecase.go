package usecase

import (
	"github.com/kafugen/ocwcentral/domain/usecase/dto"
)

type SubjectUsecase interface {
	GetById(id string) (*dto.SubjectDTO, error)
	GetByIds(ids []string) ([]*dto.SubjectDTO, error)
	GetBySearchParameter(title string, faculty string, academicField string) ([]*dto.SubjectDTO, error)
}
