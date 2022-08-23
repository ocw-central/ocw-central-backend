package model

func (c *Chapter) Id() ChapterId {
	return c.id
}
func (c *Chapter) StartAt() int {
	return c.startAt
}
func (c *Chapter) Topic() string {
	return c.topic
}
func (c *Chapter) ThumbnailLink() string {
	return c.thumbnailLink
}
