package model

import (
	"time"

	"github.com/oklog/ulid"
)

type ChapterId struct {
	id ulid.ULID `desc:"ID"`
}
type Chapter struct {
	chapterId     ChapterId `desc:"ID"`
	startAt       time.Time `desc:"チャプター開始時間"`
	topic         string    `desc:"チャプタータイトル"`
	thumbnailLink string    `desc:"サムネイルリンク"`
}
