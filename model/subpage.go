package model

import (
	"github.com/oklog/ulid"
)

type SubpageId ulid.ULID

type Subpage struct {
	id      SubpageId `desc:"ID"`
	content string    `desc:"内容テキスト"`
}
