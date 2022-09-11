package model

import (
	"github.com/oklog/ulid"
)

type ResourceId ulid.ULID

// NewResourceId generate a new ResourceId from a string or []byte representation of a ULID.
func NewResourceId[T string | []byte](ulidExp T) (*ResourceId, error) {
	var scannedULID ulid.ULID
	err := scannedULID.Scan(ulidExp)
	if err != nil {
		return nil, err
	}
	resourceId := ResourceId(scannedULID)
	return &resourceId, nil
}

func (s ResourceId) String() string {
	return ulid.ULID(s).String()
}

func (s ResourceId) ByteSlice() []byte {
	bytes := [16]byte(s)
	return bytes[:]
}

type Resource struct {
	id          ResourceId `desc:"ID"`
	title       string     `desc:"名前"`
	ordering    int        `desc:"表示順"`
	description string     `desc:"説明"`
	link        string     `desc:"リソースリンク"`
}

func NewResourceFromRepository(
	id ResourceId,
	title string,
	ordering int,
	description string,
	link string,
) *Resource {
	return &Resource{
		id:          id,
		title:       title,
		ordering:    ordering,
		description: description,
		link:        link,
	}
}
