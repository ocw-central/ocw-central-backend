package model

import (
	"github.com/oklog/ulid"
)

type PdfId ulid.ULID

type Pdf struct {
	id   PdfId  `desc:"ID"`
	link string `desc:"PDFリンク"`
}
