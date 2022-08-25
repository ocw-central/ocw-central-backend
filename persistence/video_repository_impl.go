package persistence

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/kafugen/ocwcentral/model"
	"github.com/kafugen/ocwcentral/persistence/dto"
	"github.com/kafugen/ocwcentral/utils"
)

type VideoRepositoryImpl struct {
	db *sqlx.DB
}

func NewVideoRepositoryImpl(db *sqlx.DB) VideoRepositoryImpl {
	return VideoRepositoryImpl{db}
}

func (vR VideoRepositoryImpl) GetByIds(ids []model.VideoId) ([]*model.Video, error) {
	videoIdBytes := make([]interface{}, len(ids))
	for i, id := range ids {
		videoIdBytes[i] = id.ByteSlice()
	}

	videoSQL := `
		SELECT
			videos.id,
			subject_id,
			title,
			faculty,
			ordering,
			link,
			lectured_on,
			video_length,
			language,
			chapters.id AS chapter_id,
			start_at,
			topic,
			thumbnail_link
		FROM videos
		LEFT JOIN chapters
		ON videos.id = chapters.video_id
		WHERE videos.id IN (` + utils.GetQuestionMarkStrs(len(ids)) + `)
		ORDER BY videos.ordering, chapters.start_at
	`
	var videoChapterDTOs []dto.VideoChapterDTO
	if err := vR.db.Select(&videoChapterDTOs, videoSQL, videoIdBytes...); err != nil {
		return nil, fmt.Errorf("failed on select to `videos` table: %w", err)
	}

	videos := make([]*model.Video, len(ids))

	rowIndex := 0
	for ordering, videoId := range ids {
		chapters := make([]model.Chapter, 0, 10)
		for rowIndex < len(videoChapterDTOs) && ordering == *videoChapterDTOs[rowIndex].Ordering {
			videoChapterDTO := videoChapterDTOs[rowIndex]

			if videoChapterDTO.ChapterId != nil {
				chapterId, err := model.NewChapterId(*videoChapterDTO.ChapterId)
				if err != nil {
					return nil, fmt.Errorf("failed to create `chapterId`: %w", err)
				}

				chapters = append(chapters, *model.NewChapterFromRepository(
					*chapterId,
					*videoChapterDTO.StartAt,
					*videoChapterDTO.Topic,
					*videoChapterDTO.ThumbnailLink,
				))
			}

			rowIndex++
		}

		videos[ordering] = model.NewVideoFromRepository(
			videoId,
			*videoChapterDTOs[ordering].Title,
			*videoChapterDTOs[ordering].Ordering,
			*videoChapterDTOs[ordering].Link,
			chapters,
			utils.ConvertNilToZeroValue(videoChapterDTOs[ordering].Faculty),
			utils.ConvertNilToZeroValue(videoChapterDTOs[ordering].LecturedOn),
			time.Duration(*videoChapterDTOs[ordering].VideoLength*int(time.Second)),
			utils.ConvertNilToZeroValue(videoChapterDTOs[ordering].Language),
		)
	}
	return videos, nil
}
