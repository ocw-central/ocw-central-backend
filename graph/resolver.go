package graph

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"github.com/kafugen/ocwcentral/domain/usecase"
)

type Resolver struct {
	sbU usecase.SubjectUsecase
	vU  usecase.VideoUsecase
	rU  usecase.ResourceUsecase
	slU usecase.SyllabusUsecase
	afU usecase.AcademicFieldUsecase
}

func NewResolver(
	sbU usecase.SubjectUsecase,
	vU usecase.VideoUsecase,
	rU usecase.ResourceUsecase,
	slU usecase.SyllabusUsecase,
	afU usecase.AcademicFieldUsecase,
) Resolver {
	return Resolver{
		sbU: sbU,
		vU:  vU,
		rU:  rU,
		slU: slU,
		afU: afU,
	}
}
