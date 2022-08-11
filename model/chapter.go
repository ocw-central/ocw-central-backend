package model

import (
	"time"
)

type Chapter struct {
	startAt   time.Time `desc:"チャプター開始時間"`
	topic     string    `desc:"チャプタータイトル"`
	thumbnail string    `desc:"サムネイル"`
}
