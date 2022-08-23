//go:build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/kafugen/ocwcentral/domain/repository"
	"github.com/kafugen/ocwcentral/domain/usecase"
	"github.com/kafugen/ocwcentral/env"
	"github.com/kafugen/ocwcentral/graph"
	"github.com/kafugen/ocwcentral/interactor"
	"github.com/kafugen/ocwcentral/persistence"
)

func InitializeResolver() graph.Resolver {
	wire.Build(
		graph.NewResolver,

		interactor.NewSubjectInteractor,
		interactor.NewVideoInteractor,
		interactor.NewResourceInteractor,
		interactor.NewSyllabusInteractor,
		wire.Bind(new(usecase.SubjectUsecase), new(interactor.SubjectInteractor)),
		wire.Bind(new(usecase.VideoUsecase), new(interactor.VideoInteractor)),
		wire.Bind(new(usecase.ResourceUsecase), new(interactor.ResourceInteractor)),
		wire.Bind(new(usecase.SyllabusUsecase), new(interactor.SyllabusInteractor)),

		persistence.NewSubjectRepositoryImpl,
		persistence.NewVideoRepositoryImpl,
		persistence.NewResourceRepositoryImpl,
		persistence.NewSyllabusRepositoryImpl,
		wire.Bind(new(repository.SubjectRepository), new(persistence.SubjectRepositoryImpl)),
		wire.Bind(new(repository.VideoRepository), new(persistence.VideoRepositoryImpl)),
		wire.Bind(new(repository.ResourceRepository), new(persistence.ResourceRepositoryImpl)),
		wire.Bind(new(repository.SyllabusRepository), new(persistence.SyllabusRepositoryImpl)),

		persistence.NewDB,
		env.NewEnvConfig,
	)
	return graph.Resolver{}
}
