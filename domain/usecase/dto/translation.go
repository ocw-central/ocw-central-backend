package dto

import (
	"github.com/kafugen/ocwcentral/model"
)

type TranslationDTO struct {
	ID           string `json:"id"`
	LanguageCode string `json:"languageCode"`
	Translation  string `json:"translation"`
}

func NewTranslationDTO(translation *model.Translation) *TranslationDTO {
	return &TranslationDTO{
		ID:           translation.Id().String(),
		LanguageCode: translation.LanguageCode(),
		Translation:  translation.Translation(),
	}
}
