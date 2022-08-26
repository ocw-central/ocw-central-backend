package model

import (
	"time"

	"github.com/oklog/ulid"
)

type VideoId ulid.ULID

// NewVideoId generate a new VideoId from a string or []byte representation of a ULID.
func NewVideoId[T string | []byte](ulidExp T) (*VideoId, error) {
	var scannedULID ulid.ULID
	err := scannedULID.Scan(ulidExp)
	if err != nil {
		return nil, err
	}
	videoId := VideoId(scannedULID)
	return &videoId, nil
}

func (s VideoId) String() string {
	return ulid.ULID(s).String()
}

func (s VideoId) ByteSlice() []byte {
	bytes := [16]byte(s)
	return bytes[:]
}

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

func NewVideoFromRepository(
	id VideoId,
	title string,
	ordering int,
	link string,
	chapters []Chapter,
	faculty string,
	lecturedOn time.Time,
	videoLength time.Duration,
	language string,
) *Video {
	return &Video{
		id, title, ordering, link, chapters, faculty,
		lecturedOn, videoLength, language,
	}
}
