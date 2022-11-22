package dto

type TranslationDTO struct {
	Id           *[]byte `db:"id"`
	LanguageCode *string `db:"language_code"`
	Translation  *string `db:"translation"`
}
