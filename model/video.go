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
	facultyIds  []FacultyId   `desc:"FacultyIds"`
	lecturedOn  time.Time     `desc:"講義日"`
	videoLength time.Duration `desc:"動画の長さ"`
	language    string        `desc:"使用言語"`
}
