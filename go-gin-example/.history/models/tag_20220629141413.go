package models

type Tag struct {
	Model

	Name      string `json:"name"`
	CreatedBy string `json:"created_by"`
}
