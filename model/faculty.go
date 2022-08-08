package model

import (
	"github.com/oklog/ulid"
)

type FacultyId struct {
	id ulid.ULID `desc:"ID"`
}

type Faculty struct {
	id         FacultyId `desc:"ID"`
	name       string    `desc:"名前"`
	department string    `desc:"所属部局名"`
	rank       string    `desc:"職位"`
}
