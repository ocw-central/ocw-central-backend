package persistence

import (
	"bytes"
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

// GetByIds returns videos with the given ids.
// This function is expected to get videos of one subject.
func (vR *VideoRepositoryImpl) GetByIds(ids []model.VideoId) ([]*model.Video, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	videoIdBytes := make([][]byte, len(ids))
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
		WHERE videos.id IN (?)
		ORDER BY videos.ordering, chapters.start_at
	`
	query, args, err := sqlx.In(videoSQL, videoIdBytes)
	if err != nil {
		return nil, fmt.Errorf("failed on expand `In` statement: %w", err)
	}

	var videoChapterDTOs []dto.VideoChapterDTO
	if err := vR.db.Select(&videoChapterDTOs, query, args...); err != nil {
		return nil, fmt.Errorf("failed on select to `videos` table: %w", err)
	}

	videos, err := vR.getVideosFromDTOs(videoChapterDTOs)
	if err != nil {
		return nil, fmt.Errorf("failed to get videos from DTOs: %w", err)
	}
	return videos, nil
}

// getVideosFromDTOs returns videos from the given videoChapterDTOs.
// videoChapterDTOs need not be sorted by id or ordering,
// but chapters of the same video must be contiguous.
func (vR *VideoRepositoryImpl) getVideosFromDTOs(videoChapterDTOs []dto.VideoChapterDTO) ([]*model.Video, error) {
	rowIndex := 0

	// the number of video is expected to be smaller than 20,
	// because this function is expected to be called with videos of one subject.
	videos := make([]*model.Video, 0, 20)
	for rowIndex < len(videoChapterDTOs) {
		videoChapterDTO := videoChapterDTOs[rowIndex]

		chapters, err := getChapters(videoChapterDTOs[rowIndex:])
		if err != nil {
			return nil, fmt.Errorf("failed to get chapters (rowIndex: %v): %w", rowIndex, err)
		}

		translations, err := vR.getTranslations(videoChapterDTO.Id)
		if err != nil {
			return nil, fmt.Errorf("failed to get translations (rowIndex: %v): %w", rowIndex, err)
		}

		videoId, err := model.NewVideoId(*videoChapterDTO.Id)
		if err != nil {
			return nil, fmt.Errorf("failed to create `videoId`: %w", err)
		}

		video := model.NewVideoFromRepository(
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
			translations,
		)
		videos = append(videos, video)

		if len(chapters) == 0 {
			rowIndex++
		} else {
			rowIndex += len(chapters)
		}
	}
	return videos, nil
}

func (vR *VideoRepositoryImpl) getTranslations(videoId *[]byte) ([]model.Translation, error) {
	sql := `
		SELECT
			id,
			language_code,
			translation
		FROM translations
		WHERE video_id = ?
	`

	var translationDTOs []dto.TranslationDTO
	if err := vR.db.Select(&translationDTOs, sql, videoId); err != nil {
		return nil, fmt.Errorf("failed on select to `translation` table: %w", err)
	}

	if len(translationDTOs) == 0 {
		return nil, nil
	}

	tms := make([]model.Translation, len(translationDTOs))
	for i, td := range translationDTOs {
		translationId, err := model.NewTranslationId(*td.Id)
		if err != nil {
			return nil, fmt.Errorf("failed to create `translationId`: %w", err)
		}

		tms[i] = *model.NewTranslationFromRepository(
			*translationId,
			*td.LanguageCode,
			*td.Translation,
		)
	}
	return tms, nil
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

func (vR *VideoRepositoryImpl) GetBySearchParameter(title string, faculty string) ([]*model.Video, error) {
	if title == "" && faculty == "" {
		return nil, nil
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
	`

	parameters := map[string]interface{}{"title": "%" + title + "%", "faculty": "%" + faculty + "%"}
	if title != "" && faculty != "" {
		videoSQL += "WHERE title LIKE :title AND faculty LIKE :faculty\n"
	} else if title != "" {
		videoSQL += "WHERE title LIKE :title\n"
		delete(parameters, "faculty")
	} else if faculty != "" {
		videoSQL += "WHERE faculty LIKE :faculty\n"
		delete(parameters, "title")
	}
	videoSQL += "ORDER BY videos.id, chapters.start_at"

	stmt, err := vR.db.PrepareNamed(videoSQL)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}

	var videoChapterDTOs []dto.VideoChapterDTO
	if err := stmt.Select(&videoChapterDTOs, parameters); err != nil {
		return nil, fmt.Errorf("failed on select to `videos` table: %w", err)
	}

	videos, err := vR.getVideosFromDTOs(videoChapterDTOs)
	if err != nil {
		return nil, fmt.Errorf("failed to get videos from DTOs: %w", err)
	}
	return videos, nil
}
