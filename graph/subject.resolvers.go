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
	videoDTOs, err := r.vU.GetByIds(obj.VideoIds)
	if err != nil {
		return nil, fmt.Errorf("failed on executing `GetByIds` func of VideoUsecase: %w", err)
	}

	videos := make([]*model.Video, len(videoDTOs))
	for i, videoDTO := range videoDTOs {
		chapters := make([]model.Chapter, len(videoDTO.Chapters))
		for i, chapter := range videoDTO.Chapters {
			chapters[i] = model.Chapter(chapter)
		}
		videos[i] = &model.Video{
			ID:          videoDTO.ID,
			Title:       videoDTO.Title,
			Link:        videoDTO.Link,
			Chapters:    chapters,
			Faculty:     videoDTO.Faculty,
			LecturedOn:  videoDTO.LecturedOn,
			VideoLength: videoDTO.VideoLength,
			Language:    videoDTO.Language,
		}
	}
	return videos, nil
}

// Resources is the resolver for the resources field.
func (r *subjectResolver) Resources(ctx context.Context, obj *model.Subject) ([]*model.Resource, error) {
	panic(fmt.Errorf("not implemented: Resources - resources"))
}

// RelatedSubjects is the resolver for the relatedSubjects field.
func (r *subjectResolver) RelatedSubjects(ctx context.Context, obj *model.Subject) ([]*model.RelatedSubject, error) {
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
	panic(fmt.Errorf("not implemented: Syllabus - syllabus"))
}

// Subject returns generated.SubjectResolver implementation.
func (r *Resolver) Subject() generated.SubjectResolver { return &subjectResolver{r} }

type subjectResolver struct{ *Resolver }
