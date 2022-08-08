package model

import (
	"github.com/oklog/ulid"
)

type PdfId struct {
	id ulid.ULID `desc:"ID"`
}

type Pdf struct {
	id				 PdfId `desc:"ID"`
	link			 string `desc:"PDFリンク"`
}