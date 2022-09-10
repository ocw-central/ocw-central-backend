package dto

import (
	"github.com/kafugen/ocwcentral/model"
)

type SubpageDTO struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

func NewSubpageDTO(subpage *model.Subpage) *SubpageDTO {
	return &SubpageDTO{
		ID:      subpage.Id().String(),
		Content: subpage.Content(),
	}
}
