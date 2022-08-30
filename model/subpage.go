package model

import (
	"github.com/oklog/ulid"
)

type SubpageId ulid.ULID

// NewSubpageId generate a new SubpageId from a string or []byte representation of a ULID.
func NewSubpageId[T string | []byte](ulidExp T) (*SubpageId, error) {
	var scannedULID ulid.ULID
	err := scannedULID.Scan(ulidExp)
	if err != nil {
		return nil, err
	}
	subpageId := SubpageId(scannedULID)
	return &subpageId, nil
}

func (s SubpageId) String() string {
	return ulid.ULID(s).String()
}

type Subpage struct {
	id      SubpageId `desc:"ID"`
	content string    `desc:"ページ内容"`
}
