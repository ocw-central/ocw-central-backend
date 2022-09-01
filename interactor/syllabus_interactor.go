package interactor

import (
	"fmt"

	"github.com/kafugen/ocwcentral/domain/repository"
	"github.com/kafugen/ocwcentral/domain/usecase/dto"
	"github.com/kafugen/ocwcentral/model"
)

type SyllabusInteractor struct {
	sR repository.SyllabusRepository
}

func NewSyllabusInteractor(sR repository.SyllabusRepository) SyllabusInteractor {
	return SyllabusInteractor{sR}
}

func (sI SyllabusInteractor) GetById(id string) (*dto.SyllabusDTO, error) {

	syllabusId, err := model.NewSyllabusId(id)
	if err != nil {
		return nil, fmt.Errorf("failed to create `SubjectId` struct: %w", err)
	}

	syllabus, err := sI.sR.GetById(*syllabusId)
	if err != nil {
		return nil, fmt.Errorf("failed on executing `GetById` of SyllabusRepository: %w", err)
	}

	subpageDTOs := make([]dto.SubpageDTO, len(syllabus.Subpages()))
	for i, subpage := range syllabus.Subpages() {
		subpageDTOs[i] = dto.SubpageDTO{
			ID:      subpage.Id().String(),
			Content: subpage.Content(),
		}
	}

	SyllabusDTO := &dto.SyllabusDTO{
		ID:                syllabus.Id().String(),
		Faculty:           syllabus.Faculty(),
		Language:          syllabus.Language(),
		SubjectNumbering:  syllabus.SubjectNumbering(),
		AcademicYear:      syllabus.AcademicYear(),
		Semester:          syllabus.Semester(),
		NumCredit:         syllabus.NumCredit(),
		CourseFormat:      syllabus.CourseFormat(),
		AssignedGrade:     syllabus.AssignedGrade(),
		TargetedAudience:  syllabus.TargetedAudience(),
		CourseDayPeriod:   syllabus.CourseDayPeriod(),
		Outline:           syllabus.Outline(),
		Objective:         syllabus.Objective(),
		LessonPlan:        syllabus.LessonPlan(),
		GradingMethod:     syllabus.GradingMethod(),
		CourseRequirement: syllabus.CourseRequirement(),
		OutClassLearning:  syllabus.OutClassLearning(),
		Reference:         syllabus.Reference(),
		Remark:            syllabus.Remark(),
		Subpages:          subpageDTOs,
	}
	return SyllabusDTO, nil
}
