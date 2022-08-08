package model

import (
	"github.com/oklog/ulid"
	"time"
)

type VideoId struct {
	id ulid.ULID `desc:"ID"`
}

type Video struct {
	id				 VideoId `desc:"ID"`
	title			 string `desc:"タイトル"`
	link			 string `desc:"リンク"`
	chapters		 []*Chapter `desc:"チャプターs"`
	facultyIds		 []int64 `desc:"FacultyIds"`
	lecturedOn time.Time `desc:"講義日"`
	videoLength time.Duration `desc:"動画の長さ"`
	Language		 string `desc:"使用言語"`
}
