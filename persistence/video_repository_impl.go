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
	if len(ids) == 0 {
		return nil, nil
	}

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

	rowIndex := 0
	videos := make([]*model.Video, len(ids))
	for ordering := 0; ordering < len(ids); ordering++ {
		videoChapterDTO := videoChapterDTOs[rowIndex]

		chapters, err := getChaptersByOrdering(ordering, videoChapterDTOs[rowIndex:])
		if err != nil {
			return nil, fmt.Errorf("failed to get chapters (rowIndex: %v, ordering: %v): %w", rowIndex, ordering, err)
		}

		videoId, err := model.NewVideoId(*videoChapterDTO.Id)
		if err != nil {
			return nil, fmt.Errorf("failed to create `videoId`: %w", err)
		}

		videos[ordering] = model.NewVideoFromRepository(
			*videoId,
			*videoChapterDTO.Title,
			*videoChapterDTO.Ordering,
			*videoChapterDTO.Link,
			chapters,
			utils.ConvertNilToZeroValue(videoChapterDTO.Faculty),
			utils.ConvertNilToZeroValue(videoChapterDTO.LecturedOn),
			time.Duration(*videoChapterDTO.VideoLength*int(time.Second)),
			utils.ConvertNilToZeroValue(videoChapterDTO.Language),
		)

		rowIndex += len(chapters)
	}
	return videos, nil
}

// getChaptersByOrdering returns chapters of the video with the given ordering.
func getChaptersByOrdering(ordering int, videoChapterDTOs []dto.VideoChapterDTO) ([]model.Chapter, error) {
	rowIndex := 0

	// the number of chapter is exptected to be smaller than 10.
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
	return chapters, nil
}
