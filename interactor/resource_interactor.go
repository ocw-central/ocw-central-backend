package interactor

import (
	"fmt"

	"github.com/kafugen/ocwcentral/domain/repository"
	"github.com/kafugen/ocwcentral/domain/usecase/dto"
	"github.com/kafugen/ocwcentral/model"
)

type ResourceInteractor struct {
	sR repository.ResourceRepository
}

func NewResourceInteractor(sR repository.ResourceRepository) *ResourceInteractor {
	return &ResourceInteractor{sR}
}

func (sI *ResourceInteractor) GetByIds(ids []string) ([]*dto.ResourceDTO, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	resourceIds := make([]model.ResourceId, len(ids))
	for i, id := range ids {
		resourceId, err := model.NewResourceId(id)
		resourceIds[i] = *resourceId
		if err != nil {
			return nil, fmt.Errorf("failed on create `ResourceId` struct: %w", err)
		}
	}

	resources, err := sI.sR.GetByIds(resourceIds)
	if err != nil {
		return nil, fmt.Errorf("failed on executing `GetByIds` of ResourceRepository: %w", err)
	}

	resourceDTOs := make([]*dto.ResourceDTO, len(resources))
	for i, resource := range resources {
		resourceDTOs[i] = &dto.ResourceDTO{
			ID:          resource.Id().String(),
			Title:       resource.Title(),
			Ordering:    resource.Ordering(),
			Description: resource.Description(),
			Link:        resource.Link(),
		}
	}
	return resourceDTOs, nil
}
