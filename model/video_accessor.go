package model

import (
	"time"
)

func (v *Video) Id() VideoId {
	return v.id
}
func (v *Video) Title() string {
	return v.title
}
func (v *Video) Ordering() int {
	return v.ordering
}
func (v *Video) Link() string {
	return v.link
}
func (v *Video) Chapters() []Chapter {
	return v.chapters
}
func (v *Video) Faculty() string {
	return v.faculty
}
func (v *Video) LecturedOn() time.Time {
	return v.lecturedOn
}
func (v *Video) VideoLength() time.Duration {
	return v.videoLength
}
func (v *Video) Language() string {
	return v.language
}
