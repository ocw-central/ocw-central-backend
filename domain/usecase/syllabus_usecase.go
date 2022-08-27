package usecase

import (
	"github.com/kafugen/ocwcentral/domain/usecase/dto"
)

type SyllabusUsecase interface {
	GetById(id []string) (*dto.SyllabusDTO, error)
}
