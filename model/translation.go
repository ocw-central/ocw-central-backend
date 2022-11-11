package model

import (
	"github.com/oklog/ulid"
)

type TranslationId ulid.ULID

// NewTranslationId generate a new TranslationId from a string or []byte representation of a ULID.
func NewTranslationId[T string | []byte](ulidExp T) (*TranslationId, error) {
	var scannedULID ulid.ULID
	err := scannedULID.Scan(ulidExp)
	if err != nil {
		return nil, err
	}
	translationId := TranslationId(scannedULID)
	return &translationId, nil
}

func (s TranslationId) String() string {
	return ulid.ULID(s).String()
}

func (s TranslationId) ByteSlice() []byte {
	bytes := [16]byte(s)
	return bytes[:]
}

type Translation struct {
	id           TranslationId `desc:"ID"`
	languageCode string        `desc:"言語コード"`
	translation  string        `desc:"説明"`
}

func NewTranslationFromRepository(
	id TranslationId,
	languageCode string,
	translation string,
) *Translation {
	return &Translation{
		id:           id,
		languageCode: languageCode,
		translation:  translation,
	}
}
