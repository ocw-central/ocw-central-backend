package repository

import (
	"github.com/kafugen/ocwcentral/model"
)

type SyllabusRepository interface {
	GetById(id model.SyllabusId) (*model.Syllabus, error)
}
