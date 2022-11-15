package model

func (t *Translation) Id() TranslationId {
	return t.id
}

func (t *Translation) LanguageCode() string {
	return t.languageCode
}

func (t *Translation) Translation() string {
	return t.translation
}
