package model

import (
	"github.com/oklog/ulid"
)

type SubpageId struct {
	id ulid.ULID `desc:"ID"`
}

type Subpage struct {
	id      SubpageId `desc:"ID"`
	content string    `desc:"内容テキスト"`
}
