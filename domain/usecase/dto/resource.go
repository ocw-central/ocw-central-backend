package dto

import (
	"github.com/kafugen/ocwcentral/model"
)

type ResourceDTO struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Ordering    int    `json:"Ordering"`
	Description string `json:"Description"`
	Link        string `json:"link"`
}

func NewResourceDTO(resource *model.Resource) *ResourceDTO {
	return &ResourceDTO{
		ID:          resource.Id().String(),
		Title:       resource.Title(),
		Ordering:    resource.Ordering(),
		Description: resource.Description(),
		Link:        resource.Link(),
	}
}
