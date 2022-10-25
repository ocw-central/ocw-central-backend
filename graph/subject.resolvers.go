package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/kafugen/ocwcentral/graph/generated"
	"github.com/kafugen/ocwcentral/graph/model"
)

// Videos is the resolver for the videos field.
func (r *subjectResolver) Videos(ctx context.Context, obj *model.Subject) ([]*model.Video, error) {
	if len(obj.VideoIds) == 0 {
		return nil, nil
	}

	videoDTOs, err := r.vU.GetByIds(obj.VideoIds)
	if err != nil {
		return nil, fmt.Errorf("failed on executing `GetByIds` func of VideoUsecase: %w", err)
	}

	videos := make([]*model.Video, len(videoDTOs))
	for i, videoDTO := range videoDTOs {
		videos[i] = model.NewVideo(videoDTO)
	}
	return videos, nil
}

// Resources is the resolver for the resources field.
func (r *subjectResolver) Resources(ctx context.Context, obj *model.Subject) ([]*model.Resource, error) {
	if len(obj.ResourceIds) == 0 {
		return nil, nil
	}

	resouceDTOs, err := r.rU.GetByIds(obj.ResourceIds)
	if err != nil {
		return nil, fmt.Errorf("failed on executing `GetByIds` func of ResourceUsecase: %w", err)
	}

	resources := make([]*model.Resource, len(resouceDTOs))
	for i, resouceDTO := range resouceDTOs {
		resources[i] = &model.Resource{
			ID:          resouceDTO.ID,
			Title:       resouceDTO.Title,
			Ordering:    resouceDTO.Ordering,
			Description: resouceDTO.Description,
			Link:        resouceDTO.Link,
		}
	}
	return resources, nil
}

// RelatedSubjects is the resolver for the relatedSubjects field.
func (r *subjectResolver) RelatedSubjects(ctx context.Context, obj *model.Subject) ([]*model.RelatedSubject, error) {
	if len(obj.RelatedSubjectIds) == 0 {
		return nil, nil
	}

	subjects, err := r.sbU.GetByIds(obj.RelatedSubjectIds)
	if err != nil {
		return nil, fmt.Errorf("failed on executing `GetById` func of SubjectUsecase: %w", err)
	}

	relatedSubjects := make([]*model.RelatedSubject, len(subjects))
	for i, subject := range subjects {
		relatedSubject := model.RelatedSubject(*subject)
		relatedSubjects[i] = &relatedSubject
	}
	return relatedSubjects, nil
}

// Syllabus is the resolver for the syllabus field.
func (r *subjectResolver) Syllabus(ctx context.Context, obj *model.Subject) (*model.Syllabus, error) {
	if len(obj.SyllabusId) == 0 {
		return nil, nil
	}

	syllabusDTO, err := r.slU.GetById(obj.SyllabusId)
	if err != nil {
		return nil, fmt.Errorf("failed on executing `GetByIds` func of SyllabusUsecase: %w", err)
	}

	subpages := make([]model.Subpage, len(syllabusDTO.Subpages))
	for i, subpage := range syllabusDTO.Subpages {
		subpages[i] = model.Subpage(subpage)
	}
	syllabus := &model.Syllabus{
		ID:                syllabusDTO.ID,
		Faculty:           syllabusDTO.Faculty,
		Language:          syllabusDTO.Language,
		SubjectNumbering:  syllabusDTO.SubjectNumbering,
		AcademicYear:      syllabusDTO.AcademicYear,
		Semester:          syllabusDTO.Semester,
		NumCredit:         syllabusDTO.NumCredit,
		CourseFormat:      syllabusDTO.CourseFormat,
		AssignedGrade:     syllabusDTO.AssignedGrade,
		TargetedAudience:  syllabusDTO.TargetedAudience,
		CourseDayPeriod:   syllabusDTO.CourseDayPeriod,
		Outline:           syllabusDTO.Outline,
		Objective:         syllabusDTO.Objective,
		LessonPlan:        syllabusDTO.LessonPlan,
		GradingMethod:     syllabusDTO.GradingMethod,
		CourseRequirement: syllabusDTO.CourseRequirement,
		OutClassLearning:  syllabusDTO.OutClassLearning,
		Reference:         syllabusDTO.Reference,
		Remark:            syllabusDTO.Remark,
		Subpages:          subpages,
	}

	return syllabus, nil
}

// Subject returns generated.SubjectResolver implementation.
func (r *Resolver) Subject() generated.SubjectResolver { return &subjectResolver{r} }

type subjectResolver struct{ *Resolver }
