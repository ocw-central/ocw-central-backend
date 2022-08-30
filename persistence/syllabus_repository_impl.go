package persistence

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/kafugen/ocwcentral/model"
	"github.com/kafugen/ocwcentral/persistence/dto"
	"github.com/kafugen/ocwcentral/utils"
)

type SyllabusRepositoryImpl struct {
	db *sqlx.DB
}

func NewSyllabusRepositoryImpl(db *sqlx.DB) SyllabusRepositoryImpl {
	return SyllabusRepositoryImpl{db}
}

func (sR SyllabusRepositoryImpl) GetByIds(ids []model.SyllabusId) ([]*model.Syllabus, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	syllabusIdBytes := make([]interface{}, len(ids))
	for i, id := range ids {
		syllabusIdBytes[i] = id.ByteSlice()
	}

	syllabusSQL := `
		SELECT
			syllabuses.id,
			subject_id,
			faculty,
			language,
			subject_numbering,
			academic_year,
			semester,
			num_credit,
			course_format,
			assigned_grade,
			targetted_audience,
			course_day_period,
			outline,
			objective,
			lesson_plan,
			grading_method,
			course_requirement,
			outclass_learning,
			reference,
			remark,
			subpages.id AS subpages_id,
			link
			content
		FROM syllabuses
		LEFT JOIN subpages 
		ON syllabuses.subjevt_id = subpages.subject_id
		WHERE syllabuses.id = (` + utils.GetQuestionMarkStrs(len(ids)) + `)
	`

	var syllabusSubpageDTOs []dto.SyllabusSubpageDTO
	if err := sR.db.Select(&syllabusSubpageDTOs, syllabusSQL, syllabusIdBytes); err != nil {
		return nil, fmt.Errorf("failed on select to `syllabuses` table: %w", err)
	}

	syllabuses := make([]*model.Syllabus, len(ids))
	for i, syllabusSubpageDTO := range syllabusSubpageDTOs {

		if syllabusSubpageDTO.SubpageId != nil {

			subpages := make([]model.Subpage, len(syllabusSubpageDTO.SubpageId))
			for subpageIndex := 0; subpageIndex < len(syllabusSubpageDTO.SubpageId); subpageIndex++ {

				subpageId, err := model.NewSubpageId(*syllabusSubpageDTO.SubpageId)
				if err != nil {
					return nil, fmt.Errorf("failed on creating subpageId: %w", err)
				}

				subpages[subpageIndex] = *model.NewSubpageFromRepository(
					*subpageId,
					utils.ConvertNilToZeroValue(syllabusSubpageDTO.Content),
				)
			}

			syllabusId, err := model.NewSyllabusId(*syllabusSubpageDTO.Id)
			if err != nil {
				return nil, fmt.Errorf("failed on creating syllabusId: %w", err)
			}

			syllabuses[i] = model.NewSyllabusFromRepository(
				*syllabusId,
				utils.ConvertNilToZeroValue(syllabusSubpageDTO.Faculty),
				utils.ConvertNilToZeroValue(syllabusSubpageDTO.Language),
				utils.ConvertNilToZeroValue(syllabusSubpageDTO.SubjectNumbering),
				utils.ConvertNilToZeroValue(syllabusSubpageDTO.AcademicYear),
				utils.ConvertNilToZeroValue(syllabusSubpageDTO.Semester),
				utils.ConvertNilToZeroValue(syllabusSubpageDTO.NumCredit),
				utils.ConvertNilToZeroValue(syllabusSubpageDTO.CourseFormat),
				utils.ConvertNilToZeroValue(syllabusSubpageDTO.AssignedGrade),
				utils.ConvertNilToZeroValue(syllabusSubpageDTO.TargettedAudience),
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
		}
	}
	return syllabuses, nil
}
