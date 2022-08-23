package model

import (
	"time"

	"github.com/oklog/ulid"
)

type VideoId ulid.ULID

type Video struct {
	id          VideoId       `desc:"ID"`
	title       string        `desc:"タイトル"`
	ordering    int           `desc:"順番"`
	link        string        `desc:"リンク"`
	chapters    []Chapter     `desc:"チャプターs"`
	faculty     string        `desc:"教員の氏名と所属職位"`
	lecturedOn  time.Time     `desc:"講義日"`
	videoLength time.Duration `desc:"動画の長さ"`
	language    string        `desc:"使用言語"`
}