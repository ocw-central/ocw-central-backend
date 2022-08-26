package persistence

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/kafugen/ocwcentral/model"
	"github.com/kafugen/ocwcentral/persistence/dto"
	"github.com/kafugen/ocwcentral/utils"
)

type SubjectRepositoryImpl struct {
	db *sqlx.DB
}

func NewSubjectRepositoryImpl(db *sqlx.DB) SubjectRepositoryImpl {
	return SubjectRepositoryImpl{db}
}

func (sR SubjectRepositoryImpl) GetById(id model.SubjectId) (*model.Subject, error) {
	subjectIdBytes := id.ByteSlice()

	subjectDTO := dto.SubjectDTO{}
	subjectSQL := `
		SELECT
			subjects.id,
			category,
			title,
			location,
			department,
			first_held_on,
			subjects.faculty,
			subjects.language,
			free_description,
			series,
			academic_field,
			syllabuses.id AS syllabus_id,
			thumbnail_link
		FROM subjects
		LEFT JOIN syllabuses
		ON subjects.id = syllabuses.subject_id
		WHERE subjects.id = ?
	`
	if err := sR.db.Get(&subjectDTO, subjectSQL, subjectIdBytes); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("specified subject not found: %w", err)
		}
		return nil, fmt.Errorf("failed on select to `subjects` table: %w", err)
	}

	videoIdDTOs := []struct{ Id *[]byte }{}
	resourceIdDTOs := []struct{ Id *[]byte }{}
	relatedSubjectIdDTOs := []struct{ Id *[]byte }{}

	videoIdSQL := "SELECT id FROM videos WHERE subject_id = ?"
	resourceIdSQL := "SELECT id FROM resources WHERE subject_id = ?"
	relatedSubjectIdSQL := "SELECT related_subject_id AS id FROM subject_related_subjects WHERE subject_id = ?"

	if err := sR.db.Select(&videoIdDTOs, videoIdSQL, subjectIdBytes); err != nil {
		return nil, fmt.Errorf("failed on select to `videos` table: %w", err)
	}
	if err := sR.db.Select(&resourceIdDTOs, resourceIdSQL, subjectIdBytes); err != nil {
		return nil, fmt.Errorf("failed on select to `resources` table: %w", err)
	}
	if err := sR.db.Select(&relatedSubjectIdDTOs, relatedSubjectIdSQL, subjectIdBytes); err != nil {
		return nil, fmt.Errorf("failed on select to `subject_related_subjects` table: %w", err)
	}

	videoIds := make([]model.VideoId, len(videoIdDTOs))
	resourceIds := make([]model.ResourceId, len(resourceIdDTOs))
	relatedSubjectIds := make([]model.SubjectId, len(relatedSubjectIdDTOs))

	for i, videoIdDTO := range videoIdDTOs {
		videoId, err := model.NewVideoId(*videoIdDTO.Id)
		if err != nil {
			return nil, fmt.Errorf("failed to create `videoId`: %w", err)
		}
		videoIds[i] = *videoId
	}
	for i, resourceIdDTO := range resourceIdDTOs {
		resourceId, err := model.NewResourceId(*resourceIdDTO.Id)
		if err != nil {
			return nil, fmt.Errorf("failed to create `resourceId`: %w", err)
		}
		resourceIds[i] = *resourceId
	}
	for i, relatedSubjectIdDTO := range relatedSubjectIdDTOs {
		relatedSubjectId, err := model.NewSubjectId(*relatedSubjectIdDTO.Id)
		if err != nil {
			return nil, fmt.Errorf("failed to create `relatedSubjectId`: %w", err)
		}
		relatedSubjectIds[i] = *relatedSubjectId
	}

	var syllabusId *model.SyllabusId
	if subjectDTO.SyllabusId != nil {
		var err error
		syllabusId, err = model.NewSyllabusId(*subjectDTO.SyllabusId)
		if err != nil {
			return nil, fmt.Errorf("failed to create `syllabusId`: %w", err)
		}
	} else {
		syllabusId = nil
	}

	subject := model.NewSubjectFromRepository(
		id,
		utils.ConvertNilToZeroValue(subjectDTO.Category),
		utils.ConvertNilToZeroValue(subjectDTO.Title),
		videoIds,
		utils.ConvertNilToZeroValue(subjectDTO.Location),
		resourceIds,
		relatedSubjectIds,
		utils.ConvertNilToZeroValue(subjectDTO.Department),
		utils.ConvertNilToZeroValue(subjectDTO.FirstHeldOn),
		utils.ConvertNilToZeroValue(subjectDTO.Faculty),
		utils.ConvertNilToZeroValue(subjectDTO.Language),
		utils.ConvertNilToZeroValue(subjectDTO.FreeDescription),
		syllabusId,
		utils.ConvertNilToZeroValue(subjectDTO.Series),
		utils.ConvertNilToZeroValue(subjectDTO.AcademicField),
		utils.ConvertNilToZeroValue(subjectDTO.ThumbnailLink),
	)

	return subject, nil
}

func (sR SubjectRepositoryImpl) GetByIds(id []model.SubjectId) ([]*model.Subject, error) {
	panic("not implemented")
}
