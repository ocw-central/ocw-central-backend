package model

import (
	"time"

	"github.com/oklog/ulid"
)

type ChapterId ulid.ULID

type Chapter struct {
	id            ChapterId `desc:"ID"`
	startAt       time.Time `desc:"チャプター開始時間"`
	topic         string    `desc:"チャプタータイトル"`
	thumbnailLink string    `desc:"サムネイルリンク"`
}
