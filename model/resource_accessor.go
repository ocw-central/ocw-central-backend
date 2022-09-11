package model

func (r *Resource) Id() ResourceId {
	return r.id
}
func (r *Resource) Title() string {
	return r.title
}
func (r *Resource) Ordering() int {
	return r.ordering
}
func (r *Resource) Description() string {
	return r.description
}
func (r *Resource) Link() string {
	return r.link
}
