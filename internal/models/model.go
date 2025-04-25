package models

type PageMetadata struct {
	OrderDir string `json:"order_dir"` // asc || desc

	Page int `json:"page"`
	Size int `json:"size"`
}
