package persistence

import (
	"bytes"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/kafugen/ocwcentral/model"
	"github.com/kafugen/ocwcentral/persistence/dto"
	"github.com/kafugen/ocwcentral/utils"
)

type SyllabusRepositoryImpl struct {
	db *sqlx.DB
}

func NewSyllabusRepositoryImpl(db *sqlx.DB) *SyllabusRepositoryImpl {
	return &SyllabusRepositoryImpl{db}
}

func (sR *SyllabusRepositoryImpl) GetByIds(ids []model.SyllabusId) ([]*model.Syllabus, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	syllabusIdBytes := make([][]byte, len(ids))
	for i, id := range ids {
		syllabusIdBytes[i] = id.ByteSlice()
	}

	syllabusSQL := `
		SELECT
			syllabuses.id,
			faculty,
			language,
			subject_numbering,
			academic_year,
			semester,
			num_credit,
			course_format,
			assigned_grade,
			targeted_audience,
			course_day_period,
			outline,
			objective,
			lesson_plan,
			grading_method,
			course_requirement,
			outclass_learning,
			reference,
			remark,
			subpages.id AS subpage_id,
			content
		FROM syllabuses
		LEFT JOIN subpages
		ON syllabuses.subject_id = subpages.subject_id
		WHERE syllabuses.id in (?)
		ORDER BY syllabuses.id, subpages.id
	`
	query, args, err := sqlx.In(syllabusSQL, syllabusIdBytes)
	if err != nil {
		return nil, fmt.Errorf("failed on expand `In` statement: %w", err)
	}

	var syllabusSubpageDTOs []dto.SyllabusSubpageDTO
	if err := sR.db.Select(&syllabusSubpageDTOs, query, args...); err != nil {
		return nil, fmt.Errorf("failed on select to `syllabuses` table: %w", err)
	}

	rowIndex := 0

	syllabuses := make([]*model.Syllabus, len(ids))
	for syllabusIndex := 0; syllabusIndex < len(syllabuses); syllabusIndex++ {
		syllabusSubpageDTO := syllabusSubpageDTOs[rowIndex]

		syllabusId, err := model.NewSyllabusId(*syllabusSubpageDTO.Id)
		if err != nil {
			return nil, fmt.Errorf("failed to create `syllabusId`: %w", err)
		}

		subpages, err := getSubpages(syllabusSubpageDTOs[rowIndex:])
		if err != nil {
			return nil, fmt.Errorf("failed to get subpages (rowIndex: %v): %w", rowIndex, err)
		}

		syllabuses[syllabusIndex] = model.NewSyllabusFromRepository(
			*syllabusId,
			utils.ConvertNilToZeroValue(syllabusSubpageDTO.Faculty),
			utils.ConvertNilToZeroValue(syllabusSubpageDTO.Language),
			utils.ConvertNilToZeroValue(syllabusSubpageDTO.SubjectNumbering),
			utils.ConvertNilToZeroValue(syllabusSubpageDTO.AcademicYear),
			utils.ConvertNilToZeroValue(syllabusSubpageDTO.Semester),
			utils.ConvertNilToZeroValue(syllabusSubpageDTO.NumCredit),
			utils.ConvertNilToZeroValue(syllabusSubpageDTO.CourseFormat),
			utils.ConvertNilToZeroValue(syllabusSubpageDTO.AssignedGrade),
			utils.ConvertNilToZeroValue(syllabusSubpageDTO.TargetedAudience),
			utils.ConvertNilToZeroValue(syllabusSubpageDTO.CourseDayPeriod),
			utils.ConvertNilToZeroValue(syllabusSubpageDTO.Outline),
			utils.ConvertNilToZeroValue(syllabusSubpageDTO.Objective),
			utils.ConvertNilToZeroValue(syllabusSubpageDTO.LessonPlan),
			utils.ConvertNilToZeroValue(syllabusSubpageDTO.GradingMethod),
			utils.ConvertNilToZeroValue(syllabusSubpageDTO.CourseRequirement),
			utils.ConvertNilToZeroValue(syllabusSubpageDTO.OutclassLearning),
			utils.ConvertNilToZeroValue(syllabusSubpageDTO.Reference),
			utils.ConvertNilToZeroValue(syllabusSubpageDTO.Remark),
			subpages,
		)

		if len(subpages) == 0 {
			rowIndex++
		} else {
			rowIndex += len(subpages)
		}
	}
	return syllabuses, nil
}

// GetSubpages returns a list of subpages with the same syllabusId
func getSubpages(syllabusSubpageDTOs []dto.SyllabusSubpageDTO) ([]model.Subpage, error) {
	if syllabusSubpageDTOs[0].SubpageId == nil {
		return nil, nil
	}

	rowIndex := 0

	// number of subpages is expected to be less than 20
	subpages := make([]model.Subpage, 0, 20)

	for rowIndex < len(syllabusSubpageDTOs) && bytes.Equal(*syllabusSubpageDTOs[0].Id, *syllabusSubpageDTOs[rowIndex].Id) {
		syllabusSubpageDTO := syllabusSubpageDTOs[rowIndex]

		if syllabusSubpageDTO.SubpageId == nil {
			return subpages, nil
		}

		subpageId, err := model.NewSubpageId(*syllabusSubpageDTO.SubpageId)
		if err != nil {
			return nil, fmt.Errorf("failed to create `subpageId`: %w", err)
		}

		subpages = append(subpages, *model.NewSubpageFromRepository(
			*subpageId,
			utils.ConvertNilToZeroValue(syllabusSubpageDTO.Content),
		))
		rowIndex++
	}
	return subpages, nil

}

func (sR *SyllabusRepositoryImpl) GetById(id model.SyllabusId) (*model.Syllabus, error) {
	syllabuses, err := sR.GetByIds([]model.SyllabusId{id})
	if err != nil {
		return nil, fmt.Errorf("failed to get syllabus : %w", err)
	}

	if len(syllabuses) == 0 {
		return nil, nil
	} else {
		return syllabuses[0], nil
	}
}
