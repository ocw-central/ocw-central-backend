package model

import (
	"github.com/oklog/ulid"
)

type ChapterId ulid.ULID

type Chapter struct {
	id            ChapterId `desc:"ID"`
	startAt       int       `desc:"チャプター開始時間"`
	topic         string    `desc:"チャプタータイトル"`
	thumbnailLink string    `desc:"サムネイルリンク"`
}
