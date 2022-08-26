package persistence

import (
	"database/sql"
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

func (sR SyllabusRepositoryImpl) GetById(id model.SyllabusId) (*model.Syllabus, error) {
	syllabusIdBytes := id.ByteSlice()

	syllabusDTO := dto.SyllabusDTO{}
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
		FROM syllabuses
		LEFT JOIN subpages 
		ON syllabuses.subjevt_id = subpages.subject_id
		WHERE syllabuses.id = ?
	`
	var syllabusSubpageDTOs []dto.SyllabusSubpageDTO
	if err := sR.db.Select(&syllabusSubpageDTOs, syllabusSQL, syllabusIdBytes); err != nil {
		return nil, fmt.Errorf("failed on select to `syllabuses` table: %w", err)
	}

	if err := sR.db.Get(&syllabusDTO, syllabusSQL, syllabusIdBytes); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("specified syllabus not found: %w", err)
		}
		return nil, fmt.Errorf("failed on select to `syllabuses` table: %w", err)
	}

	syllabusId := *(*[16]byte)(*syllabusDTO.Id)

	syllabus := model.NewSyllabusFromRepository(
		model.SyllabusId(syllabusId),
		utils.ConvertNilToZeroValue(syllabusDTO.Faculty),
		utils.ConvertNilToZeroValue(syllabusDTO.Language),
		utils.ConvertNilToZeroValue(syllabusDTO.SubjectNumbering),
		utils.ConvertNilToZeroValue(syllabusDTO.AcademicYear),
		utils.ConvertNilToZeroValue(syllabusDTO.Semester),
		utils.ConvertNilToZeroValue(syllabusDTO.NumCredit),
		utils.ConvertNilToZeroValue(syllabusDTO.CourseFormat),
		utils.ConvertNilToZeroValue(syllabusDTO.AssignedGrade),
		utils.ConvertNilToZeroValue(syllabusDTO.TargettedAudience),
		utils.ConvertNilToZeroValue(syllabusDTO.CourseDayPeriod),
		utils.ConvertNilToZeroValue(syllabusDTO.Outline),
		utils.ConvertNilToZeroValue(syllabusDTO.Objective),
		utils.ConvertNilToZeroValue(syllabusDTO.LessonPlan),
		utils.ConvertNilToZeroValue(syllabusDTO.GradingMethod),
		utils.ConvertNilToZeroValue(syllabusDTO.CourseRequirement),
		utils.ConvertNilToZeroValue(syllabusDTO.OutclassLearning),
		utils.ConvertNilToZeroValue(syllabusDTO.Reference),
		utils.ConvertNilToZeroValue(syllabusDTO.Remark),
		subpageIds,
	)

	return syllabus, nil
}

func (sR SubjectRepositoryImpl) GetByIds(id []model.SyllabusId) ([]*model.Syllabus, error) {
	panic("not implemented")
}
