package model

import (
	"github.com/oklog/ulid"
)

type ChapterId ulid.ULID

// NewChapterId generate a new ChapterId from a string or []byte representation of a ULID.
func NewChapterId[T string | []byte](ulidExp T) (*ChapterId, error) {
	var scannedULID ulid.ULID
	err := scannedULID.Scan(ulidExp)
	if err != nil {
		return nil, err
	}
	chapterId := ChapterId(scannedULID)
	return &chapterId, nil
}

func (s ChapterId) String() string {
	return ulid.ULID(s).String()
}

type Chapter struct {
	id            ChapterId `desc:"ID"`
	startAt       int       `desc:"チャプター開始時間"`
	topic         string    `desc:"チャプタータイトル"`
	thumbnailLink string    `desc:"サムネイルリンク"`
}
