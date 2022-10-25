package dto

import (
	"github.com/kafugen/ocwcentral/model"
)

type SubjectWithSpecifiedVideosDTO struct {
	Subject SubjectDTO
	Videos  []VideoDTO
}

func NewSubjectWithSpecifiedVideosDTO(
	subject *model.Subject,
	videos []*model.Video,
) *SubjectWithSpecifiedVideosDTO {
	videoDTOs := make([]VideoDTO, len(videos))
	for i, video := range videos {
		videoDTOs[i] = *NewVideoDTO(video)
	}
	return &SubjectWithSpecifiedVideosDTO{
		Subject: *NewSubjectDTO(subject),
		Videos:  videoDTOs,
	}
}
