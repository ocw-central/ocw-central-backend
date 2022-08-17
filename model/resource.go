package model

import (
	"github.com/oklog/ulid"
)

type ResourceId struct {
	id ulid.ULID `desc:"ID"`
}

type Resource struct {
	id          ResourceId `desc:"ID"`
	title       string     `desc:"名前"`
	ordering    int        `desc:"表示順"`
	description string     `desc:"説明"`
	link        string     `desc:"リソースリンク"`
}
