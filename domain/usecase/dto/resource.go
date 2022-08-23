package dto

type ResourceDTO struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Ordering    int    `json:"Ordering"`
	Description string `json:"Description"`
	Link        string `json:"link"`
}
