package persistence

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/kafugen/ocwcentral/model"
	"github.com/kafugen/ocwcentral/persistence/dto"
	"github.com/kafugen/ocwcentral/utils"
)

type ResourceRepositoryImpl struct {
	db *sqlx.DB
}

func NewResourceRepositoryImpl(db *sqlx.DB) ResourceRepositoryImpl {
	return ResourceRepositoryImpl{db}
}

func (vR ResourceRepositoryImpl) GetByIds(ids []model.ResourceId) ([]*model.Resource, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	resourceIdBytes := make([]interface{}, len(ids))
	for i, id := range ids {
		resourceIdBytes[i] = id.ByteSlice()
	}

	resourceSQL := `
		SELECT
			resources.id,
			title,
			description,
			ordering,
			link
		FROM resources
		WHERE resources.id IN (` + utils.GetQuestionMarkStrs(len(ids)) + `)
		ORDER BY ordering
	`

	var resourceDTOs []dto.ResourceDTO
	if err := vR.db.Select(&resourceDTOs, resourceSQL, resourceIdBytes...); err != nil {
		return nil, fmt.Errorf("failed on select to `resources` table: %w", err)
	}

	resources := make([]*model.Resource, len(ids))
	for rowIndex := 0; rowIndex < len(ids); rowIndex++ {
		resourceDTO := resourceDTOs[rowIndex]

		resourceId, err := model.NewResourceId(*resourceDTO.Id)
		if err != nil {
			return nil, fmt.Errorf("failed to create `resourceId`: %w", err)
		}

		resources[rowIndex] = model.NewResourceFromRepository(
			*resourceId,
			utils.ConvertNilToZeroValue(resourceDTO.Title),
			*resourceDTO.Ordering,
			utils.ConvertNilToZeroValue(resourceDTO.Description),
			utils.ConvertNilToZeroValue(resourceDTO.Link),
		)
	}
	return resources, nil
}
