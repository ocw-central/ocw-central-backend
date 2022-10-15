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

func NewVideoRepositoryImpl(db *sqlx.DB) *VideoRepositoryImpl {
	return &VideoRepositoryImpl{db}
}

func (vR *VideoRepositoryImpl) GetByIds(ids []model.VideoId) ([]*model.Video, error) {
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
			thumbnail_link,
			transcription
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

		chapters, err := getChapters(videoChapterDTOs[rowIndex:])
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
			utils.ConvertNilToZeroValue(videoChapterDTO.Transcription),
		)

		if len(chapters) == 0 {
			rowIndex++
		} else {
			rowIndex += len(chapters)
		}
	}
	return videos, nil
}

// getChapters returns chapters of the video with the given ordering.
// videoChapterDTOs need not be sorted by id or ordering,
// but chapters of the same video must be contiguous.
// The video containing the chapter to retrieve must come at the top of the array.
func getChapters(videoChapterDTOs []dto.VideoChapterDTO) ([]model.Chapter, error) {
	videoId := videoChapterDTOs[0].Id
	rowIndex := 0

	// the number of chapter is expected to be smaller than 10.
	chapters := make([]model.Chapter, 0, 10)
	for rowIndex < len(videoChapterDTOs) && bytes.Equal(*videoId, *videoChapterDTOs[rowIndex].Id) {
		videoChapterDTO := videoChapterDTOs[rowIndex]

		if videoChapterDTO.ChapterId == nil {
			return chapters, nil
		}

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

		rowIndex++
	}
	return chapters, nil
}
